
/**
 * Samples for VirtualNetworkLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/VirtualNetworkLinkList.json
     */
    /**
     * Sample code: Get Private DNS Zone Virtual Network Links.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        getPrivateDNSZoneVirtualNetworkLinks(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getVirtualNetworkLinks().list("resourceGroup1", "privatelink.contoso.com", null,
            com.azure.core.util.Context.NONE);
    }
}
