
/**
 * Samples for NetworkSecurityPerimeterAccessRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspAccessRuleList.json
     */
    /**
     * Sample code: NspAccessRulesList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspAccessRulesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterAccessRules().list("rg1", "nsp1",
            "profile1", null, null, com.azure.core.util.Context.NONE);
    }
}
