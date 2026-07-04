
/**
 * Samples for NetworkManagers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerList.json
     */
    /**
     * Sample code: List Network Manager.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listNetworkManager(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkManagers().listByResourceGroup("rg1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
