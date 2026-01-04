
import com.azure.resourcemanager.network.models.TroubleshootingParameters;

/**
 * Samples for NetworkWatchers GetTroubleshooting.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherTroubleshootGet
     * .json
     */
    /**
     * Sample code: Get troubleshooting.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTroubleshooting(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().getTroubleshooting("rg1", "nw1",
            new TroubleshootingParameters()
                .withTargetResourceId(
                    "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1")
                .withStorageId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/st1")
                .withStoragePath("https://st1.blob.core.windows.net/cn1"),
            com.azure.core.util.Context.NONE);
    }
}
