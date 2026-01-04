
/**
 * Samples for NetworkSecurityPerimeterAccessRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspAccessRuleGet.json
     */
    /**
     * Sample code: NspAccessRuleGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspAccessRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterAccessRules().getWithResponse("rg1",
            "nsp1", "profile1", "accessRule1", com.azure.core.util.Context.NONE);
    }
}
