
/**
 * Samples for NetworkVirtualAppliances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceDelete.json
     */
    /**
     * Sample code: Delete NetworkVirtualAppliance.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteNetworkVirtualAppliance(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().delete("rg1", "nva", com.azure.core.util.Context.NONE);
    }
}
