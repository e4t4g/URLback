# GetStatApi

All URIs are relative to *http://localhost:8006*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getStat**](GetStatApi.md#getStat) | **GET** /linkUrl/[id] | Get stat link


<a name="getStat"></a>
# **getStat**
> Line getStat(urlStat)

Get stat link

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.GetStatApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8006");

    GetStatApi apiInstance = new GetStatApi(defaultClient);
    String urlStat = "urlStat_example"; // String | stat info
    try {
      Line result = apiInstance.getStat(urlStat);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling GetStatApi#getStat");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlStat** | **String**| stat info |

### Return type

[**Line**](Line.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**400** | Invalid link |  -  |

