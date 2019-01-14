## **getCatalogInfo**

#### **接口功能**
> 
获取制定项目的分类信息

#### **URL**
> 
[/snserver/getCatalogInfo]()

#### **支持格式**
> 
JSON

#### **请求方式**
> 
HTTP-GET

#### **请求字段**
|名称|必选|类型|父字段|说明|
|:--  |:--  |:-- |:-- |---- |
|name |Y |String | |请求的项目名11 |
|type |Y |Integer | |请求项目的类型。1：类型一；2：类型二 。|

#### **返回字段**
|名称|必选|类型|父字段|说明|
|:--  |:--  |:-- |:-- |---- |
|status |Y |Integer | |返回结果状态。0：正常；1：错误。 |
|company |Y |Company | | 所属公司名 |
|name |Y |Integer |company | 所属公司名 |
|category |Y |Integer | |所属类型 |

#### **错误码**
|返回码|说明 |
|:-- |:--- |
|4001 | 错误1 |
|4002 | 错误2 |
|4003 | 错误3 |

#### **接口示例**
地址：[http://www.api.com/index.php?name="可口可乐"&type=1](http://www.api.com/index.php?name="可口可乐"&type=1)
``` json
{
    "statue": 0,
    "company": {
    "name":"可口可乐"
    },
    "category": "饮料",
}
```


