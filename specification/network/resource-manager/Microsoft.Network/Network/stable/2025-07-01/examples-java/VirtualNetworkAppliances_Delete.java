
/**
 * Samples for VirtualNetworkAppliances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkAppliances_Delete.json
     */
    /**
     * Sample code: Delete virtual network appliance.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteVirtualNetworkAppliance(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkAppliances().delete("rg1", "test-vna",
            com.azure.core.util.Context.NONE);
    }
}
