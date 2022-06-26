import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.FlowLogInformationInner;

/** Samples for NetworkWatchers SetFlowLogConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkWatcherFlowLogConfigure.json
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
