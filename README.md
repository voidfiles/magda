# Magda

A platform for emmersing one self in the power of beauty, chance, and juxtaposition.

# What?

There are so many beautiful things in this world. Art, a quote, a picture from your loved ones. As we pass through life we are enriched when we encounter these things. But, so often they are hidden from our eye, or they are arranged so staticially that we miss out on all the new thigns being produced everyday.

Wouldn't it be great if you could emerser you self in those touch points. For them to slowly evolve around you everyday. From your house, to your phone, to the screens around you.

Magda is my attempt to build a platform for that experience.

# How?

Great question. I don't know exactly.

## Contributors

**Prerequists**: Right now, you need to have working go environment.

Then you can run: `make setup` to install required tools.

To regenerate the support tools run: `make compile`

## The Plan

- [] Create the data model
    - [] Entrys
    - [] Entitys
    - [] Collections
- [] Create an API
- [] Create mockups of some of the screens
- [] Create a storybook for the components
- [] Figure out the Authentication story
    - [] Roles
        - [] Viewer (possibly anonymous), can look at collections
        - [] User (Can add things to the system that belong to them)
        - [] Admin (Can do anything )

# Data Model

![Data Model](https://raw.githubusercontent.com/voidfiles/magda/master/gen/magda.dot.png "Output")

# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [service.proto](#service.proto)
    - [EntryRequest](#magda.EntryRequest)
    - [EntryResponse](#magda.EntryResponse)
  
  
  
    - [MagdaService](#magda.MagdaService)
  

- [entry.proto](#entry.proto)
    - [Entry](#magda.Entry)
    - [Source](#magda.Source)
  
    - [Entry.Kind](#magda.Entry.Kind)
  
  
  

- [file.proto](#file.proto)
    - [File](#magda.File)
  
  
  
  

- [entity.proto](#entity.proto)
    - [Entity](#magda.Entity)
  
    - [Entity.Kind](#magda.Entity.Kind)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service.proto



<a name="magda.EntryRequest"></a>

### EntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="magda.EntryResponse"></a>

### EntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [Entry](#magda.Entry) |  |  |





 

 

 


<a name="magda.MagdaService"></a>

### MagdaService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetEntry | [EntryRequest](#magda.EntryRequest) | [EntryResponse](#magda.EntryResponse) |  |

 



<a name="entry.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## entry.proto



<a name="magda.Entry"></a>

### Entry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique ID for entry |
| kind | [Entry.Kind](#magda.Entry.Kind) |  | kind of entry |
| source | [Source](#magda.Source) |  | Source of entry |
| titles | [string](#string) | repeated | Titles of entry, canonical is first |
| files | [File](#magda.File) | repeated | Files for entry, canonical is first |
| description | [string](#string) |  | Description of entry for display |
| creators | [Entity](#magda.Entity) | repeated | Creators of entry (Author/Painter, etc) |
| published_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Work created at data, may be empty |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Entry created in magada at |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Entry updated at |






<a name="magda.Source"></a>

### Source



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#magda.Entity) |  | Source entity |
| url | [string](#string) |  | URL of source, usually a website specific to entry |
| titles | [string](#string) | repeated | Titles of source, canonical is first |





 


<a name="magda.Entry.Kind"></a>

### Entry.Kind


| Name | Number | Description |
| ---- | ------ | ----------- |
| QUOTE | 0 |  |
| IMAGE | 1 |  |


 

 

 



<a name="file.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## file.proto



<a name="magda.File"></a>

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



<a name="magda.Entity"></a>

### Entity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| kind | [Entity.Kind](#magda.Entity.Kind) |  |  |
| names | [string](#string) | repeated |  |
| urls | [string](#string) | repeated |  |
| description | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| files | [File](#magda.File) | repeated |  |





 


<a name="magda.Entity.Kind"></a>

### Entity.Kind


| Name | Number | Description |
| ---- | ------ | ----------- |
| PERSON | 0 |  |
| ORGANIZATION | 1 |  |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

