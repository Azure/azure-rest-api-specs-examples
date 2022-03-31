Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ProductApi CheckEntityExists. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadProductApi.json
     */
    /**
     * Sample code: ApiManagementHeadProductApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadProductApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productApis()
            .checkEntityExistsWithResponse(
                "rg1", "apimService1", "5931a75ae4bbd512a88c680b", "59306a29e4bbd510dc24e5f9", Context.NONE);
    }
}
```
