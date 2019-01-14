# APIGen
### generate api request and response from md

#### mdtableparse -h
>Usage of mdtableparse:
>  -d string
>    	the path of data to generator the > java (default "~")
>  -i string
>    	the name of the interface
>  -o string
>    	the path of data to output 
> (default "~")
>  -p string
>    	the package of the java

#### for example
> mdtableparse -d ~/Desktop/template.md -o ~/Desktop -p org.springcat -i WWW

#### result:
> packagePath: org.springcat
> interaceName: WWW
> dataPath: /Users/springcat/Desktop/template.md
> /Users/springcat/Desktop/WWWRequest.java gen success
> /Users/springcat/Desktop/WWWResponse.java gen success

#### cat WWWRequest.java
```Java
package org.springcat;

import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@NoArgsConstructor
@Builder
@Data
/**
*
* @author springcat
*/
public class WWWRequest{

	/**
	 * must: Y
	 * 请求的项目名11
	 *
	*/
	private String name;

	/**
	 * must: Y
	 * 请求项目的类型。1：类型一；2：类型二 。
	 *
	*/
	private Integer type;

}
```

#### cat WWWResponse.java
```Java
package org.springcat;

import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@NoArgsConstructor
@Builder
@Data
/**
*
*  |4001 | 错误1 |
*  |4002 | 错误2 |
*  |4003 | 错误3 |
* @author springcat
*/
public class WWWResponse{

	/**
	 * must: Y
	 * 返回结果状态。0：正常；1：错误。
	 *
	*/
	private Integer status;

	/**
	 * must: Y
	 * 所属公司名
	 *
	*/
	private Company company;

	/**
	 * must: Y
	 * 所属类型
	 *
	*/
	private Integer category;

    @Data
    @NoArgsConstructor
    @Builder
    class Company{

        /**
        * must: Y
        * 所属公司名
        *
        */
        private Integer name;

    }

```


