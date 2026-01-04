
/**
 * Samples for SecurityUserRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerSecurityUserRuleList.json
     */
    /**
     * Sample code: List security user rules.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSecurityUserRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserRules().list("rg1", "testNetworkManager",
            "myTestSecurityConfig", "testRuleCollection", null, null, com.azure.core.util.Context.NONE);
    }
}
