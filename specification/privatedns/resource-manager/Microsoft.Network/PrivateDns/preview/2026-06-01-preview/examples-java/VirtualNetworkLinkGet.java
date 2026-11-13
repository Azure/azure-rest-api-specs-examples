
/**
 * Samples for VirtualNetworkLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/VirtualNetworkLinkGet.json
     */
    /**
     * Sample code: GET Private DNS Zone Virtual Network Link.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        gETPrivateDNSZoneVirtualNetworkLink(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getVirtualNetworkLinks().getWithResponse("resourceGroup1", "privatezone1.com",
            "virtualNetworkLink1", com.azure.core.util.Context.NONE);
    }
}
