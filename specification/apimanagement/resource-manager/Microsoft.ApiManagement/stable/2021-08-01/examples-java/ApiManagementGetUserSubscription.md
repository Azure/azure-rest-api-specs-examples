Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for UserSubscription Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetUserSubscription.json
     */
    /**
     * Sample code: ApiManagementGetUserSubscription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetUserSubscription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .userSubscriptions()
            .getWithResponse("rg1", "apimService1", "1", "5fa9b096f3df14003c070001", Context.NONE);
    }
}
```
