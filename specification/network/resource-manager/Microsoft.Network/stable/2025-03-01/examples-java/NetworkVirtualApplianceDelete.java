
/**
 * Samples for NetworkVirtualAppliances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkVirtualApplianceDelete
     * .json
     */
    /**
     * Sample code: Delete NetworkVirtualAppliance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkVirtualAppliance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkVirtualAppliances().delete("rg1", "nva",
            com.azure.core.util.Context.NONE);
    }
}
