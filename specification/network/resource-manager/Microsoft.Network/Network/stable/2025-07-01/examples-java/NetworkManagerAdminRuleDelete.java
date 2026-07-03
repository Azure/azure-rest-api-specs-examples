
/**
 * Samples for AdminRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerAdminRuleDelete.json
     */
    /**
     * Sample code: Deletes an admin rule.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletesAnAdminRule(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAdminRules().delete("rg1", "testNetworkManager", "myTestSecurityConfig",
            "testRuleCollection", "SampleAdminRule", false, com.azure.core.util.Context.NONE);
    }
}
