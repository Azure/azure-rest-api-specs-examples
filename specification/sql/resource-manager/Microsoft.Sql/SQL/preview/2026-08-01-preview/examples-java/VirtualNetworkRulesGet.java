
/**
 * Samples for VirtualNetworkRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/VirtualNetworkRulesGet.json
     */
    /**
     * Sample code: Gets a virtual network rule.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAVirtualNetworkRule(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getVirtualNetworkRules().getWithResponse("Default", "vnet-test-svr",
            "vnet-firewall-rule", com.azure.core.util.Context.NONE);
    }
}
