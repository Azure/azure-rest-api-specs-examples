
/**
 * Samples for VirtualNetworkLinks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/VirtualNetworkLinkDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone Virtual Network Link.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        dELETEPrivateDNSZoneVirtualNetworkLink(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getVirtualNetworkLinks().delete("resourceGroup1", "privatezone1.com",
            "virtualNetworkLink1", null, com.azure.core.util.Context.NONE);
    }
}
