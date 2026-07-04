
/**
 * Samples for AdminRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerAdminRuleGet.json
     */
    /**
     * Sample code: Gets security admin rule.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getsSecurityAdminRule(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAdminRules().getWithResponse("rg1", "testNetworkManager", "myTestSecurityConfig",
            "testRuleCollection", "SampleAdminRule", com.azure.core.util.Context.NONE);
    }
}
