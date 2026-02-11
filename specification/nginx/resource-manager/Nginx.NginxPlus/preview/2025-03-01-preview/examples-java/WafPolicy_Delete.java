
/**
 * Samples for WafPolicy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/WafPolicy_Delete.json
     */
    /**
     * Sample code: WafPolicy_Delete.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void wafPolicyDelete(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.wafPolicies().delete("myResourceGroup", "myDeployment", "myWafPolicy",
            com.azure.core.util.Context.NONE);
    }
}
