// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DeviceDefinition is documented here http://hl7.org/fhir/StructureDefinition/DeviceDefinition
type DeviceDefinition struct {
	Id                      *string                               `json:"id,omitempty"`
	Meta                    *Meta                                 `json:"meta,omitempty"`
	ImplicitRules           *string                               `json:"implicitRules,omitempty"`
	Language                *string                               `json:"language,omitempty"`
	Text                    *Narrative                            `json:"text,omitempty"`
	Extension               []Extension                           `json:"extension,omitempty"`
	ModifierExtension       []Extension                           `json:"modifierExtension,omitempty"`
	Identifier              []Identifier                          `json:"identifier,omitempty"`
	UdiDeviceIdentifier     []DeviceDefinitionUdiDeviceIdentifier `json:"udiDeviceIdentifier,omitempty"`
	DeviceName              []DeviceDefinitionDeviceName          `json:"deviceName,omitempty"`
	ModelNumber             *string                               `json:"modelNumber,omitempty"`
	Type                    *CodeableConcept                      `json:"type,omitempty"`
	Specialization          []DeviceDefinitionSpecialization      `json:"specialization,omitempty"`
	Version                 []string                              `json:"version,omitempty"`
	Safety                  []CodeableConcept                     `json:"safety,omitempty"`
	ShelfLifeStorage        []ProductShelfLife                    `json:"shelfLifeStorage,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic                   `json:"physicalCharacteristics,omitempty"`
	LanguageCode            []CodeableConcept                     `json:"languageCode,omitempty"`
	Capability              []DeviceDefinitionCapability          `json:"capability,omitempty"`
	Property                []DeviceDefinitionProperty            `json:"property,omitempty"`
	Owner                   *Reference                            `json:"owner,omitempty"`
	Contact                 []ContactPoint                        `json:"contact,omitempty"`
	Url                     *string                               `json:"url,omitempty"`
	OnlineInformation       *string                               `json:"onlineInformation,omitempty"`
	Note                    []Annotation                          `json:"note,omitempty"`
	Quantity                *Quantity                             `json:"quantity,omitempty"`
	ParentDevice            *Reference                            `json:"parentDevice,omitempty"`
	Material                []DeviceDefinitionMaterial            `json:"material,omitempty"`
}
type DeviceDefinitionUdiDeviceIdentifier struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DeviceIdentifier  string      `json:"deviceIdentifier"`
	Issuer            string      `json:"issuer"`
	Jurisdiction      string      `json:"jurisdiction"`
}
type DeviceDefinitionDeviceName struct {
	Id                *string        `json:"id,omitempty"`
	Extension         []Extension    `json:"extension,omitempty"`
	ModifierExtension []Extension    `json:"modifierExtension,omitempty"`
	Name              string         `json:"name"`
	Type              DeviceNameType `json:"type"`
}
type DeviceDefinitionSpecialization struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	SystemType        string      `json:"systemType"`
	Version           *string     `json:"version,omitempty"`
}
type DeviceDefinitionCapability struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	Description       []CodeableConcept `json:"description,omitempty"`
}
type DeviceDefinitionProperty struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	ValueQuantity     []Quantity        `json:"valueQuantity,omitempty"`
	ValueCode         []CodeableConcept `json:"valueCode,omitempty"`
}
type DeviceDefinitionMaterial struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	Substance           CodeableConcept `json:"substance"`
	Alternate           *bool           `json:"alternate,omitempty"`
	AllergenicIndicator *bool           `json:"allergenicIndicator,omitempty"`
}
type OtherDeviceDefinition DeviceDefinition

// MarshalJSON marshals the given DeviceDefinition as JSON into a byte slice
func (r DeviceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceDefinition: OtherDeviceDefinition(r),
		ResourceType:          "DeviceDefinition",
	})
}

// UnmarshalDeviceDefinition unmarshals a DeviceDefinition.
func UnmarshalDeviceDefinition(b []byte) (DeviceDefinition, error) {
	var deviceDefinition DeviceDefinition
	if err := json.Unmarshal(b, &deviceDefinition); err != nil {
		return deviceDefinition, err
	}
	return deviceDefinition, nil
}