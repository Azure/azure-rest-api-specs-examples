Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.apimanagement.models.ContentFormat;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiUsingImportOverrideServiceUrl.json
     */
    /**
     * Sample code: ApiManagementCreateApiUsingImportOverrideServiceUrl.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiUsingImportOverrideServiceUrl(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("apidocs")
            .withExistingService("rg1", "apimService1")
            .withValue("http://apimpimportviaurl.azurewebsites.net/api/apidocs/")
            .withFormat(ContentFormat.fromString("swagger-link"))
            .withServiceUrl("http://petstore.swagger.wordnik.com/api")
            .withPath("petstoreapi123")
            .create();
    }
}
```
