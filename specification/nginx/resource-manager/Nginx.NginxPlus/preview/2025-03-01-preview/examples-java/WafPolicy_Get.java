
/**
 * Samples for WafPolicy Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/WafPolicy_Get.json
     */
    /**
     * Sample code: WafPolicy_Get.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void wafPolicyGet(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.wafPolicies().getWithResponse("myResourceGroup", "myDeployment", "myWafPolicy",
            com.azure.core.util.Context.NONE);
    }
}
