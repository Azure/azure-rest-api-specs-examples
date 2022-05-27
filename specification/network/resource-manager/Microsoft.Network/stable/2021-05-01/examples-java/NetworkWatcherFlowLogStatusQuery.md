Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.FlowLogStatusParameters;

/** Samples for NetworkWatchers GetFlowLogStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherFlowLogStatusQuery.json
     */
    /**
     * Sample code: Get flow log status.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFlowLogStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkWatchers()
            .getFlowLogStatus(
                "rg1",
                "nw1",
                new FlowLogStatusParameters()
                    .withTargetResourceId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
                Context.NONE);
    }
}
```
