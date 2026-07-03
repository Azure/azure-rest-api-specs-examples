
/**
 * Samples for DdosCustomPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DdosCustomPolicyDelete.json
     */
    /**
     * Sample code: Delete DDoS custom policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteDDoSCustomPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDdosCustomPolicies().delete("rg1", "test-ddos-custom-policy",
            com.azure.core.util.Context.NONE);
    }
}
