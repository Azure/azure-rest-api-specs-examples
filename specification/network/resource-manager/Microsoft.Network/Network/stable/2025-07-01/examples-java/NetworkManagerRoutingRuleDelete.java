
/**
 * Samples for RoutingRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerRoutingRuleDelete.json
     */
    /**
     * Sample code: Deletes a routing rule.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletesARoutingRule(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutingRules().delete("rg1", "testNetworkManager", "myTestRoutingConfig",
            "testRuleCollection", "sampleRule", false, com.azure.core.util.Context.NONE);
    }
}
