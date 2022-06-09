```java
import com.azure.core.util.Context;

/** Samples for Subscription GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadSubscription.json
     */
    /**
     * Sample code: ApiManagementHeadSubscription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadSubscription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .subscriptions()
            .getEntityTagWithResponse("rg1", "apimService1", "5931a769d8d14f0ad8ce13b8", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
