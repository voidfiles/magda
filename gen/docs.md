# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [service.proto](#service.proto)
    - [EntryRequest](#service.EntryRequest)
    - [EntryResponse](#service.EntryResponse)
  
  
  
    - [MagdaService](#service.MagdaService)
  

- [entry.proto](#entry.proto)
    - [Entry](#service.Entry)
    - [Source](#service.Source)
  
    - [Entry.Kind](#service.Entry.Kind)
  
  
  

- [file.proto](#file.proto)
    - [File](#service.File)
  
  
  
  

- [entity.proto](#entity.proto)
    - [Entity](#service.Entity)
  
    - [Entity.Kind](#service.Entity.Kind)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service.proto



<a name="service.EntryRequest"></a>

### EntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="service.EntryResponse"></a>

### EntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [Entry](#service.Entry) |  |  |





 

 

 


<a name="service.MagdaService"></a>

### MagdaService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetEntry | [EntryRequest](#service.EntryRequest) | [EntryResponse](#service.EntryResponse) |  |

 



<a name="entry.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## entry.proto



<a name="service.Entry"></a>

### Entry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique ID for entry |
| kind | [Entry.Kind](#service.Entry.Kind) |  | kind of entry |
| source | [Source](#service.Source) |  | Source of entry |
| titles | [string](#string) | repeated | Titles of entry, canonical is first |
| files | [File](#service.File) | repeated | Files for entry, canonical is first |
| description | [string](#string) |  | Description of entry for display |
| creators | [Entity](#service.Entity) | repeated | Creators of entry (Author/Painter, etc) |
| published_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Work created at data, may be empty |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Entry created in magada at |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Entry updated at |






<a name="service.Source"></a>

### Source



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#service.Entity) |  | Source entity |
| url | [string](#string) |  | URL of source, usually a website specific to entry |
| titles | [string](#string) | repeated | Titles of source, canonical is first |





 


<a name="service.Entry.Kind"></a>

### Entry.Kind


| Name | Number | Description |
| ---- | ------ | ----------- |
| QUOTE | 0 |  |
| IMAGE | 1 |  |


 

 

 



<a name="file.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## file.proto



<a name="service.File"></a>

### File



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  |  |
| source_url | [string](#string) |  |  |
| content_type | [string](#string) |  |  |
| filename | [string](#string) |  |  |
| width | [int64](#int64) |  |  |
| hieght | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 

 

 

 



<a name="entity.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## entity.proto



<a name="service.Entity"></a>

### Entity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| kind | [Entity.Kind](#service.Entity.Kind) |  |  |
| names | [string](#string) | repeated |  |
| urls | [string](#string) | repeated |  |
| description | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| files | [File](#service.File) | repeated |  |





 


<a name="service.Entity.Kind"></a>

### Entity.Kind


| Name | Number | Description |
| ---- | ------ | ----------- |
| PERSON | 0 |  |
| ORGANIZATION | 1 |  |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

