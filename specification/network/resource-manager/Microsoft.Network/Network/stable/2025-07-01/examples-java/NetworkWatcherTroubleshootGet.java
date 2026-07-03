
import com.azure.resourcemanager.network.models.TroubleshootingParameters;

/**
 * Samples for NetworkWatchers GetTroubleshooting.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherTroubleshootGet.json
     */
    /**
     * Sample code: Get troubleshooting.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getTroubleshooting(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().getTroubleshooting("rg1", "nw1", new TroubleshootingParameters()
            .withTargetResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1")
            .withStorageId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/st1")
            .withStoragePath("https://st1.blob.core.windows.net/cn1"), com.azure.core.util.Context.NONE);
    }
}
