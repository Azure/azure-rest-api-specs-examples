
/**
 * Samples for AdminRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerAdminRuleDelete
     * .json
     */
    /**
     * Sample code: Deletes an admin rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesAnAdminRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAdminRules().delete("rg1", "testNetworkManager",
            "myTestSecurityConfig", "testRuleCollection", "SampleAdminRule", false, com.azure.core.util.Context.NONE);
    }
}
