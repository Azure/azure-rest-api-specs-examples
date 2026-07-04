
/**
 * Samples for SecurityUserRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityUserRuleList.json
     */
    /**
     * Sample code: List security user rules.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listSecurityUserRules(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserRules().list("rg1", "testNetworkManager", "myTestSecurityConfig",
            "testRuleCollection", null, null, com.azure.core.util.Context.NONE);
    }
}
