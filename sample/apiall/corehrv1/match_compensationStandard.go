// Package corehr code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/corehr/v1"
)

// GET /open-apis/corehr/v1/compensation_standards/match
func main(){
   // 创建 Client
   client := lark.NewClient("appID", "appSecret")
   // 创建请求对象
   req := larkcorehr.NewMatchCompensationStandardReqBuilder().
		UserIdType("open_id").
		
		DepartmentIdType("people_corehr_department_id").
		
		EmploymentId("7124293751317038636").
		
		ReferenceObjectApi("cpst_item").
		
		ReferenceObjectId("7156853394442044972").
		
		DepartmentId("od-53899868dd0da32292a2d809f0518c8f").
		
		WorkLocationId("7094869485965870636").
		
		CompanyId("7091599096804394540").
		
		JobFamilyId("7039313681989502508").
		
		JobLevelId("7086415175263258156").
		
		EmployeeTypeId("7039310401359775276").
		
		RecruitmentType("experienced_professionals").
		
		CpstChangeReasonId("6967639606963471117").
		
		CpstPlanId("6967639606963471118").
		
		CpstSalaryLevelId("6967639606963471119").
		
		EffectiveTime("1660924800000").
		
	   Build()
   // 发起请求
   resp,err := client.Corehr.V1.CompensationStandard.Match(context.Background(),req)


   // 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}

