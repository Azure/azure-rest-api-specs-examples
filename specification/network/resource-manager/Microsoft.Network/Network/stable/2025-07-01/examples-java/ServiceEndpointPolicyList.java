
/**
 * Samples for ServiceEndpointPolicies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceEndpointPolicyList.json
     */
    /**
     * Sample code: List resource group service endpoint policies.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listResourceGroupServiceEndpointPolicies(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceEndpointPolicies().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
