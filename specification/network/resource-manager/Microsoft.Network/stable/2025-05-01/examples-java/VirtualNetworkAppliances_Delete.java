
/**
 * Samples for VirtualNetworkAppliances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkAppliances_Delete.json
     */
    /**
     * Sample code: Delete virtual network appliance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteVirtualNetworkAppliance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkAppliances().delete("rg1", "test-vna",
            com.azure.core.util.Context.NONE);
    }
}
