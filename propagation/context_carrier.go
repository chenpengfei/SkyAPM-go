/*
 * Licensed to the OpenSkywalking under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package propagation

type CarrierItem interface {
	HeadKey() string
	HeadValue() string
	SetValue(t string)
	IsValid() bool
}

// SW3CarrierItem is the only implementation of CarrierItem for now.
// In roadmap, W3C trace context implementation may be added and actived by somehow.
type SW3CarrierItem struct {
}

func (s *SW3CarrierItem) HeadKey() string {
	return "sw3"
}

func (s *SW3CarrierItem) HeadValue() string {
	return ""
}

func (s *SW3CarrierItem) SetValue(t string) {
}

func (s *SW3CarrierItem) IsValid() bool {
	return true
}

func NewSW3CarrierItem() CarrierItem {
	item := new(SW3CarrierItem)

	return item
}

// ContextCarrier is a data carrier of tracing context,
// it holds a snapshot for across process propagation.
type ContextCarrier struct {
	items []CarrierItem
}

func (c *ContextCarrier) GetAllItems() []CarrierItem {
	return c.items
}

func NewContextCarrier() *ContextCarrier {
	carrier := ContextCarrier{items: []CarrierItem{
		NewSW3CarrierItem(),
	}}
	return &carrier
}
