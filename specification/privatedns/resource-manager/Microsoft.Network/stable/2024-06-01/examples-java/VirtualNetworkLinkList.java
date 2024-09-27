
/**
 * Samples for VirtualNetworkLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/VirtualNetworkLinkList.
     * json
     */
    /**
     * Sample code: Get Private DNS Zone Virtual Network Links.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateDNSZoneVirtualNetworkLinks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getVirtualNetworkLinks().list("resourceGroup1",
            "privatelink.contoso.com", null, com.azure.core.util.Context.NONE);
    }
}
