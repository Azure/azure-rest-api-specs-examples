
/**
 * Samples for AdminRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerAdminRuleList.json
     */
    /**
     * Sample code: List security admin rules.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listSecurityAdminRules(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAdminRules().list("rg1", "testNetworkManager", "myTestSecurityConfig",
            "testRuleCollection", null, null, com.azure.core.util.Context.NONE);
    }
}
