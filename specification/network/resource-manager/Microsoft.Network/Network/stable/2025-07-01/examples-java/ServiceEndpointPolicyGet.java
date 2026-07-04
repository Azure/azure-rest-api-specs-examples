
/**
 * Samples for ServiceEndpointPolicies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceEndpointPolicyGet.json
     */
    /**
     * Sample code: Get service endPoint Policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getServiceEndPointPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceEndpointPolicies().getByResourceGroupWithResponse("rg1",
            "testServiceEndpointPolicy", null, com.azure.core.util.Context.NONE);
    }
}
