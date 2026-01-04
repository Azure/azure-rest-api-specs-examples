
/**
 * Samples for NetworkSecurityPerimeterServiceTags List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspServiceTagsList.json
     */
    /**
     * Sample code: NSPServiceTagsList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nSPServiceTagsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterServiceTags().list("westus",
            com.azure.core.util.Context.NONE);
    }
}
