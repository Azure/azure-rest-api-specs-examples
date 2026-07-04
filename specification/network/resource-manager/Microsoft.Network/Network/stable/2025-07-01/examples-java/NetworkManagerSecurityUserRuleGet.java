
/**
 * Samples for SecurityUserRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityUserRuleGet.json
     */
    /**
     * Sample code: Gets a security user rule.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getsASecurityUserRule(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserRules().getWithResponse("rg1", "testNetworkManager",
            "myTestSecurityConfig", "testRuleCollection", "SampleUserRule", com.azure.core.util.Context.NONE);
    }
}
