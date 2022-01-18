Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for FlowLogs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherFlowLogGet.json
     */
    /**
     * Sample code: Get flow log.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFlowLog(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFlowLogs()
            .getWithResponse("rg1", "nw1", "flowLog1", Context.NONE);
    }
}
```
