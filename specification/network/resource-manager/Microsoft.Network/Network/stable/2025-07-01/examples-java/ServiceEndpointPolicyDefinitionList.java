
/**
 * Samples for ServiceEndpointPolicyDefinitions ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceEndpointPolicyDefinitionList.json
     */
    /**
     * Sample code: List service endpoint definitions in service end point policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listServiceEndpointDefinitionsInServiceEndPointPolicy(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceEndpointPolicyDefinitions().listByResourceGroup("rg1", "testPolicy",
            com.azure.core.util.Context.NONE);
    }
}
