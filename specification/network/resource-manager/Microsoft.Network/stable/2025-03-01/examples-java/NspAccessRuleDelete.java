
/**
 * Samples for NetworkSecurityPerimeterAccessRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspAccessRuleDelete.json
     */
    /**
     * Sample code: NspAccessRulesDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspAccessRulesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterAccessRules().deleteWithResponse("rg1",
            "nsp1", "profile1", "accessRule1", com.azure.core.util.Context.NONE);
    }
}
