
/**
 * Samples for SecurityUserRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityUserRuleDelete.json
     */
    /**
     * Sample code: Delete a security user rule.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteASecurityUserRule(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserRules().delete("rg1", "testNetworkManager", "myTestSecurityConfig",
            "testRuleCollection", "SampleUserRule", false, com.azure.core.util.Context.NONE);
    }
}
