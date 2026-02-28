
/**
 * Samples for NetworkSecurityPerimeterProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspProfileList.json
     */
    /**
     * Sample code: NspProfilesList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspProfilesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterProfiles().list("rg1", "nsp1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
