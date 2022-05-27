Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.FlowLogInformationInner;

/** Samples for NetworkWatchers SetFlowLogConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherFlowLogConfigure.json
     */
    /**
     * Sample code: Configure flow log.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configureFlowLog(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkWatchers()
            .setFlowLogConfiguration(
                "rg1",
                "nw1",
                new FlowLogInformationInner()
                    .withTargetResourceId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1")
                    .withStorageId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/st1")
                    .withEnabled(true),
                Context.NONE);
    }
}
```
