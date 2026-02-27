
/**
 * Samples for RoutingRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerRoutingRuleDelete.json
     */
    /**
     * Sample code: Deletes a routing rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesARoutingRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRoutingRules().delete("rg1", "testNetworkManager",
            "myTestRoutingConfig", "testRuleCollection", "sampleRule", false, com.azure.core.util.Context.NONE);
    }
}
