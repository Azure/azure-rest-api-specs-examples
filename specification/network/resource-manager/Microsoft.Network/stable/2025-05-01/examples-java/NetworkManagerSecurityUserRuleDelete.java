
/**
 * Samples for SecurityUserRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerSecurityUserRuleDelete.json
     */
    /**
     * Sample code: Delete a security user rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteASecurityUserRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserRules().delete("rg1", "testNetworkManager",
            "myTestSecurityConfig", "testRuleCollection", "SampleUserRule", false, com.azure.core.util.Context.NONE);
    }
}
