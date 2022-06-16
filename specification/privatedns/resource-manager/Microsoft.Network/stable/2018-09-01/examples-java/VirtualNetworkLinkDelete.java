import com.azure.core.util.Context;

/** Samples for VirtualNetworkLinks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/examples/VirtualNetworkLinkDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone Virtual Network Link.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dELETEPrivateDNSZoneVirtualNetworkLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .privateDnsZones()
            .manager()
            .serviceClient()
            .getVirtualNetworkLinks()
            .delete("resourceGroup1", "privatezone1.com", "virtualNetworkLink1", null, Context.NONE);
    }
}
