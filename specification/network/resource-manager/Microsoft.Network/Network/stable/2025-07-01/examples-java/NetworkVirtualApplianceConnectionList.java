
/**
 * Samples for NetworkVirtualApplianceConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceConnectionList.json
     */
    /**
     * Sample code: NetworkVirtualApplianceConnectionList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkVirtualApplianceConnectionList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualApplianceConnections().list("rg1", "nva1",
            com.azure.core.util.Context.NONE);
    }
}
