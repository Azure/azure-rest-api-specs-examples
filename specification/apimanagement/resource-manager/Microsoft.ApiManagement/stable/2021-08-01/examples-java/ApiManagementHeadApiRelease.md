Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ApiRelease GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiRelease.json
     */
    /**
     * Sample code: ApiManagementHeadApiRelease.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiRelease(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiReleases()
            .getEntityTagWithResponse("rg1", "apimService1", "a1", "5a7cb545298324c53224a799", Context.NONE);
    }
}
```
