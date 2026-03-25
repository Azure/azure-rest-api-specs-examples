
/**
 * Samples for WafPolicy Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/WafPolicy_Create.json
     */
    /**
     * Sample code: WafPolicy_Create.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void wafPolicyCreate(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.wafPolicies().define("myWafPolicy").withExistingNginxDeployment("myResourceGroup", "myDeployment")
            .create();
    }
}
