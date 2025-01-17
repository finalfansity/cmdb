package impl

import (
	"context"
	"database/sql"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"github.com/infraboard/mcube/sqlbuilder"

	"github.com/infraboard/cmdb/apps/host"
)

func (s *service) SyncHost(ctx context.Context, ins *host.Host) (
	*host.Host, error) {
	exist, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(ins.Base.Id))
	if err != nil {
		// 如果不是Not Found则直接返回
		if !exception.IsNotFoundError(err) {
			return nil, err
		}
	}

	// 检查ins已经存在 我们则需要更新ins
	if exist != nil {
		if err := s.update(ctx, ins); err != nil {
			return nil, err
		}
		return ins, nil
	}

	// 如果没有我们则直接保存
	if err := s.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) QueryHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.HostSet, error) {
	query := sqlbuilder.NewQuery(queryHostSQL)

	if req.Keywords != "" {
		query.Where("r.name LIKE ? OR r.id = ? OR r.instance_id = ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
			"%"+req.Keywords+"%",
			req.Keywords,
			req.Keywords,
			req.Keywords+"%",
			req.Keywords+"%",
		)
	}

	querySQL, args := query.Order("sync_at").Desc().Limit(req.OffSet(), uint(req.Page.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query host error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := host.NewHostSet()
	var (
		publicIPList, privateIPList, keyPairNameList, securityGroupsList string
	)
	for rows.Next() {
		ins := host.NewDefaultHost()
		base := ins.Base
		info := ins.Information
		desc := ins.Describe
		err := rows.Scan(
			&base.Id, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
			&info.Category, &info.Type, &info.Name, &info.Description,
			&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
			&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash,
			&base.SecretId, &base.Id,
			&desc.Cpu, &desc.Memory, &desc.GpuAmount, &desc.GpuSpec, &desc.OsType, &desc.OsName,
			&desc.SerialNumber, &desc.ImageId, &desc.InternetMaxBandwidthOut, &desc.InternetMaxBandwidthIn,
			&keyPairNameList, &securityGroupsList,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query host error, %s", err.Error())
		}
		info.LoadPrivateIPString(privateIPList)
		info.LoadPublicIPString(publicIPList)
		desc.LoadKeyPairNameString(keyPairNameList)
		desc.LoadSecurityGroupsString(securityGroupsList)
		set.Add(ins)
	}

	// 获取total SELECT COUNT(*) FROMT t Where ....
	countSQL, args := query.BuildCount()
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	return set, nil
}

func (s *service) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (
	*host.Host, error) {
	query := sqlbuilder.NewQuery(queryHostSQL)
	cond, val := req.Where()
	querySQL, args := query.Where(cond, val).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare describe host error, %s", err.Error())
	}
	defer queryStmt.Close()

	ins := host.NewDefaultHost()
	var (
		publicIPList, privateIPList, keyPairNameList, securityGroupsList string
	)
	base := ins.Base
	info := ins.Information
	desc := ins.Describe
	err = queryStmt.QueryRow(args...).Scan(
		&base.Id, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
		&info.Category, &info.Type, &info.Name, &info.Description,
		&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
		&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash,
		&base.SecretId, &base.Id,
		&desc.Cpu, &desc.Memory, &desc.GpuAmount, &desc.GpuSpec, &desc.OsType, &desc.OsName,
		&desc.SerialNumber, &desc.ImageId, &desc.InternetMaxBandwidthOut, &desc.InternetMaxBandwidthIn,
		&keyPairNameList, &securityGroupsList,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("%#v not found", req)
		}
		return nil, exception.NewInternalServerError("describe host error, %s", err.Error())
	}

	info.LoadPrivateIPString(privateIPList)
	info.LoadPublicIPString(publicIPList)
	desc.LoadKeyPairNameString(keyPairNameList)
	desc.LoadSecurityGroupsString(securityGroupsList)

	return ins, nil
}

func (s *service) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (
	*host.Host, error) {
	// 检测参数合法性
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update host error, %s", err)
	}

	// 查询出该条实例的数据
	ins, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PATCH:
		ins.Patch(req.UpdateHostData)
	default:
		ins.Put(req.UpdateHostData)
	}

	// 更新云商主机

	// 更新数据库
	if err := s.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) ReleaseHost(ctx context.Context, req *host.ReleaseHostRequest) (
	*host.Host, error) {
	ins, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	// 删除云商上该主机

	if err := s.delete(ctx, req); err != nil {
		return nil, err
	}

	return ins, nil
}
