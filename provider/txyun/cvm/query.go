package cvm

import (
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/utils"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

// 查看实例列表: https://cloud.tencent.com/document/api/213/15728
func (o *CVMOperater) Query(req *cvm.DescribeInstancesRequest) (*host.HostSet, error) {
	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(resp.Response.InstanceSet)
	set.Total = utils.PtrInt64(resp.Response.TotalCount)

	return set, nil
}

func NewPageQueryRequest(reqPerSecond int) *PageQueryRequest {
	return &PageQueryRequest{
		ReqPerSecond: reqPerSecond,
	}
}

type PageQueryRequest struct {
	ReqPerSecond int
}

func (o *CVMOperater) PageQuery(req *PageQueryRequest) host.Pager {
	return newPager(20, o, req.ReqPerSecond)
}
