// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type Icmpv4_bulk struct {
	Id     types.String      `tfsdk:"id"`
	Domain types.String      `tfsdk:"domain"`
	Bulk   []Icmpv4_bulkBulk `tfsdk:"bulk"`
}

type Icmpv4_bulkBulk struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	IcmpType    types.String `tfsdk:"icmp_type"`
	Code        types.Int64  `tfsdk:"code"`
	Description types.String `tfsdk:"description"`
	Overridable types.Bool   `tfsdk:"overridable"`
	ParentId    types.String `tfsdk:"parent_id"`
	ParentType  types.String `tfsdk:"parent_type"`
	TargetId    types.String `tfsdk:"target_id"`
	TargetType  types.String `tfsdk:"target_type"`
	TargetName  types.String `tfsdk:"target_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Icmpv4_bulk) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/icmpv4objects"
}

// End of section. //template:end getPath

func (data Icmpv4_bulk) toBody(ctx context.Context, state Icmpv4_bulk) string {
	body := ""
	// Don't need id field to be passed in the body
	// if data.Id.ValueString() != "" {
	// 	body, _ = sjson.Set(body, "id", data.Id.ValueString())
	// }
	if len(data.Bulk) > 0 {
		body = "[]" // Initialising the body to be an array for the bulk request
		for _, item := range data.Bulk {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.IcmpType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "icmpType", item.IcmpType.ValueString())
			}
			if !item.Code.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "code", item.Code.ValueInt64())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.Overridable.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overridable", item.Overridable.ValueBool())
			}
			if !item.ParentId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overrides.parent.id", item.ParentId.ValueString())
			}
			if !item.ParentType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overrides.parent.type", item.ParentType.ValueString())
			}
			if !item.TargetId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overrides.target.id", item.TargetId.ValueString())
			}
			if !item.TargetType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overrides.target.type", item.TargetType.ValueString())
			}
			if !item.TargetName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overrides.target.name", item.TargetName.ValueString())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *Icmpv4_bulk) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("bulk"); value.Exists() {
		data.Bulk = make([]Icmpv4_bulkBulk, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := Icmpv4_bulkBulk{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("icmpType"); cValue.Exists() {
				item.IcmpType = types.StringValue(cValue.String())
			} else {
				item.IcmpType = types.StringNull()
			}
			if cValue := v.Get("code"); cValue.Exists() {
				item.Code = types.Int64Value(cValue.Int())
			} else {
				item.Code = types.Int64Null()
			}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			} else {
				item.Description = types.StringNull()
			}
			if cValue := v.Get("overridable"); cValue.Exists() {
				item.Overridable = types.BoolValue(cValue.Bool())
			} else {
				item.Overridable = types.BoolNull()
			}
			if cValue := v.Get("overrides.parent.id"); cValue.Exists() {
				item.ParentId = types.StringValue(cValue.String())
			} else {
				item.ParentId = types.StringNull()
			}
			if cValue := v.Get("overrides.parent.type"); cValue.Exists() {
				item.ParentType = types.StringValue(cValue.String())
			} else {
				item.ParentType = types.StringNull()
			}
			if cValue := v.Get("overrides.target.id"); cValue.Exists() {
				item.TargetId = types.StringValue(cValue.String())
			} else {
				item.TargetId = types.StringNull()
			}
			if cValue := v.Get("overrides.target.type"); cValue.Exists() {
				item.TargetType = types.StringValue(cValue.String())
			} else {
				item.TargetType = types.StringNull()
			}
			if cValue := v.Get("overrides.target.name"); cValue.Exists() {
				item.TargetName = types.StringValue(cValue.String())
			} else {
				item.TargetName = types.StringNull()
			}
			data.Bulk = append(data.Bulk, item)
			return true
		})
	}
}

// End of section. //template:end fromBody