
/**
 * Samples for ServiceEndpointPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceEndpointPolicyListAll.json
     */
    /**
     * Sample code: List all service endpoint policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllServiceEndpointPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceEndpointPolicies().list(com.azure.core.util.Context.NONE);
    }
}
