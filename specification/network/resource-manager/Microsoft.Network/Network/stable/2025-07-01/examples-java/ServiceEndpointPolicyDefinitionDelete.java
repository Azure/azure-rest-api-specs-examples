
/**
 * Samples for ServiceEndpointPolicyDefinitions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceEndpointPolicyDefinitionDelete.json
     */
    /**
     * Sample code: Delete service endpoint policy definitions from service endpoint policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteServiceEndpointPolicyDefinitionsFromServiceEndpointPolicy(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceEndpointPolicyDefinitions().delete("rg1", "testPolicy", "testDefinition",
            com.azure.core.util.Context.NONE);
    }
}
