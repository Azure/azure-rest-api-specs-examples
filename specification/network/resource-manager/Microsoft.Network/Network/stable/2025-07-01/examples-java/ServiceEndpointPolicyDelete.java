
/**
 * Samples for ServiceEndpointPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceEndpointPolicyDelete.json
     */
    /**
     * Sample code: Delete service endpoint policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteServiceEndpointPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceEndpointPolicies().delete("rg1", "serviceEndpointPolicy1",
            com.azure.core.util.Context.NONE);
    }
}
