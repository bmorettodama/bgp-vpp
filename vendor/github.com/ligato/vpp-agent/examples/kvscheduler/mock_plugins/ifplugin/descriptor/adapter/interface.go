// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/examples/kvscheduler/mock_plugins/ifplugin/model"
	"github.com/ligato/vpp-agent/pkg/idxvpp"
)

////////// type-safe key-value pair with metadata //////////

type InterfaceKVWithMetadata struct {
	Key      string
	Value    *mock_interfaces.Interface
	Metadata *idxvpp.OnlyIndex
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type InterfaceDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *mock_interfaces.Interface) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *mock_interfaces.Interface) error
	Create               func(key string, value *mock_interfaces.Interface) (metadata *idxvpp.OnlyIndex, err error)
	Delete               func(key string, value *mock_interfaces.Interface, metadata *idxvpp.OnlyIndex) error
	Update               func(key string, oldValue, newValue *mock_interfaces.Interface, oldMetadata *idxvpp.OnlyIndex) (newMetadata *idxvpp.OnlyIndex, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *mock_interfaces.Interface, metadata *idxvpp.OnlyIndex) bool
	Retrieve             func(correlate []InterfaceKVWithMetadata) ([]InterfaceKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *mock_interfaces.Interface) []KeyValuePair
	Dependencies         func(key string, value *mock_interfaces.Interface) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type InterfaceDescriptorAdapter struct {
	descriptor *InterfaceDescriptor
}

func NewInterfaceDescriptor(typedDescriptor *InterfaceDescriptor) *KVDescriptor {
	adapter := &InterfaceDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *InterfaceDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castInterfaceValue(key, oldValue)
	typedNewValue, err2 := castInterfaceValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *InterfaceDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *InterfaceDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *InterfaceDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castInterfaceValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castInterfaceValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castInterfaceMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *InterfaceDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castInterfaceMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *InterfaceDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castInterfaceValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castInterfaceValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castInterfaceMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *InterfaceDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []InterfaceKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castInterfaceValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castInterfaceMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			InterfaceKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *InterfaceDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *InterfaceDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castInterfaceValue(key string, value proto.Message) (*mock_interfaces.Interface, error) {
	typedValue, ok := value.(*mock_interfaces.Interface)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castInterfaceMetadata(key string, metadata Metadata) (*idxvpp.OnlyIndex, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(*idxvpp.OnlyIndex)
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
