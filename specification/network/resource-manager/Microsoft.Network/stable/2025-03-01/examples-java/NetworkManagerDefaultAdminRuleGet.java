
/**
 * Samples for AdminRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerDefaultAdminRuleGet.json
     */
    /**
     * Sample code: Gets security default admin rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsSecurityDefaultAdminRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAdminRules().getWithResponse("rg1", "testNetworkManager",
            "myTestSecurityConfig", "testRuleCollection", "SampleDefaultAdminRule", com.azure.core.util.Context.NONE);
    }
}
