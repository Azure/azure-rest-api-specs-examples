
/**
 * Samples for DefaultWafPolicy List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/DefaultWafPolicy_List.json
     */
    /**
     * Sample code: DefaultWafPolicy_List.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void defaultWafPolicyList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.defaultWafPolicies().listWithResponse("myResourceGroup", "myDeployment",
            com.azure.core.util.Context.NONE);
    }
}
