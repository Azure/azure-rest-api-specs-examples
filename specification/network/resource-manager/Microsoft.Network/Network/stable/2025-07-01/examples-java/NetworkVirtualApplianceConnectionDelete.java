
/**
 * Samples for NetworkVirtualApplianceConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceConnectionDelete.json
     */
    /**
     * Sample code: NetworkVirtualApplianceConnectionDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        networkVirtualApplianceConnectionDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualApplianceConnections().delete("rg1", "nva1", "connection1",
            com.azure.core.util.Context.NONE);
    }
}
