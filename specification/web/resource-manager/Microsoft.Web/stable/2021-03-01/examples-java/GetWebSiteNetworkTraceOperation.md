Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WebApps GetNetworkTraceOperationSlotV2. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebSiteNetworkTraceOperation.json
     */
    /**
     * Sample code: Get the current status of a network trace operation for a site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheCurrentStatusOfANetworkTraceOperationForASite(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getNetworkTraceOperationSlotV2WithResponse(
                "testrg123", "SampleApp", "c291433b-53ad-4c49-8cae-0a293eae1c6d", "Production", Context.NONE);
    }
}
```
