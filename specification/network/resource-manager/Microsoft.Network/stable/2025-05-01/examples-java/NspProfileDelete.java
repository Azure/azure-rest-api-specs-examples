
/**
 * Samples for NetworkSecurityPerimeterProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspProfileDelete.json
     */
    /**
     * Sample code: NspProfilesDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspProfilesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterProfiles().deleteWithResponse("rg1",
            "nsp1", "profile1", com.azure.core.util.Context.NONE);
    }
}
