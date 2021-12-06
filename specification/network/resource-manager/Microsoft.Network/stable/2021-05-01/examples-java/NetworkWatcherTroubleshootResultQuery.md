Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.QueryTroubleshootingParameters;

/** Samples for NetworkWatchers GetTroubleshootingResult. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherTroubleshootResultQuery.json
     */
    /**
     * Sample code: Get troubleshoot result.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTroubleshootResult(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkWatchers()
            .getTroubleshootingResult(
                "rg1",
                "nw1",
                new QueryTroubleshootingParameters()
                    .withTargetResourceId(
                        "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
                Context.NONE);
    }
}
```
