Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.apimanagement.models.ApiCreateOrUpdatePropertiesWsdlSelector;
import com.azure.resourcemanager.apimanagement.models.ContentFormat;
import com.azure.resourcemanager.apimanagement.models.SoapApiType;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateSoapPassThroughApiUsingWsdlImport.json
     */
    /**
     * Sample code: ApiManagementCreateSoapPassThroughApiUsingWsdlImport.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateSoapPassThroughApiUsingWsdlImport(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("soapApi")
            .withExistingService("rg1", "apimService1")
            .withValue("http://www.webservicex.net/CurrencyConvertor.asmx?WSDL")
            .withFormat(ContentFormat.WSDL_LINK)
            .withWsdlSelector(
                new ApiCreateOrUpdatePropertiesWsdlSelector()
                    .withWsdlServiceName("CurrencyConvertor")
                    .withWsdlEndpointName("CurrencyConvertorSoap"))
            .withSoapApiType(SoapApiType.SOAP)
            .withPath("currency")
            .create();
    }
}
```
