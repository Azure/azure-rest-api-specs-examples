Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WebApps StartNetworkTraceSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StartWebSiteNetworkTraceOperation.json
     */
    /**
     * Sample code: Start a new network trace operation for a site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startANewNetworkTraceOperationForASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .startNetworkTraceSlot("testrg123", "SampleApp", "Production", 60, null, null, Context.NONE);
    }
}
```
