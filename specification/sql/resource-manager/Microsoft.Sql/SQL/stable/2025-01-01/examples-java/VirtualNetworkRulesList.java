
/**
 * Samples for VirtualNetworkRules ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/VirtualNetworkRulesList.json
     */
    /**
     * Sample code: List virtual network rules.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listVirtualNetworkRules(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getVirtualNetworkRules().listByServer("Default", "vnet-test-svr",
            com.azure.core.util.Context.NONE);
    }
}
