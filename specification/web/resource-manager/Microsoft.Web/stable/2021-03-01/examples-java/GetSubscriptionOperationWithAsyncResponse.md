Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Global GetSubscriptionOperationWithAsyncResponse. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetSubscriptionOperationWithAsyncResponse.json
     */
    /**
     * Sample code: Gets an operation in a subscription and given region.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAnOperationInASubscriptionAndGivenRegion(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getGlobals()
            .getSubscriptionOperationWithAsyncResponseWithResponse(
                "West US", "34adfa4f-cedf-4dc0-ba29-b6d1a69ab5d5", Context.NONE);
    }
}
```
