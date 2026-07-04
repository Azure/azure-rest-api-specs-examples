
/**
 * Samples for DdosCustomPolicies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DdosCustomPolicyGet.json
     */
    /**
     * Sample code: Get DDoS custom policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getDDoSCustomPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDdosCustomPolicies().getByResourceGroupWithResponse("rg1", "test-ddos-custom-policy",
            com.azure.core.util.Context.NONE);
    }
}
