
/**
 * Samples for ServiceEndpointPolicyDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceEndpointPolicyDefinitionGet.json
     */
    /**
     * Sample code: Get service endpoint definition in service endpoint policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getServiceEndpointDefinitionInServiceEndpointPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceEndpointPolicyDefinitions().getWithResponse("rg1", "testPolicy",
            "testDefinition", com.azure.core.util.Context.NONE);
    }
}
