
/**
 * Samples for AdminRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerDefaultAdminRuleGet.json
     */
    /**
     * Sample code: Gets security default admin rule.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getsSecurityDefaultAdminRule(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAdminRules().getWithResponse("rg1", "testNetworkManager", "myTestSecurityConfig",
            "testRuleCollection", "SampleDefaultAdminRule", com.azure.core.util.Context.NONE);
    }
}
