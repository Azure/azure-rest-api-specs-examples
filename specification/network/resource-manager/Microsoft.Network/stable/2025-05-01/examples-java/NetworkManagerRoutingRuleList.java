
/**
 * Samples for RoutingRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerRoutingRuleList
     * .json
     */
    /**
     * Sample code: List routing rules.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoutingRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRoutingRules().list("rg1", "testNetworkManager",
            "myTestRoutingConfig", "testRuleCollection", null, null, com.azure.core.util.Context.NONE);
    }
}
