Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ProductSubscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListProductSubscriptions.json
     */
    /**
     * Sample code: ApiManagementListProductSubscriptions.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListProductSubscriptions(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productSubscriptions()
            .list("rg1", "apimService1", "5600b57e7e8880006a060002", null, null, null, Context.NONE);
    }
}
```
