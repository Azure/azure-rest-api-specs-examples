
/**
 * Samples for WafPolicy List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/WafPolicy_List.json
     */
    /**
     * Sample code: WafPolicy_List.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void wafPolicyList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.wafPolicies().list("myResourceGroup", "myDeployment", com.azure.core.util.Context.NONE);
    }
}
