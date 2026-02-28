
/**
 * Samples for NetworkSecurityPerimeterProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspProfileGet.json
     */
    /**
     * Sample code: NspProfilesGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspProfilesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterProfiles().getWithResponse("rg1", "nsp1",
            "profile1", com.azure.core.util.Context.NONE);
    }
}
