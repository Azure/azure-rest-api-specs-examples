
/**
 * Samples for ServiceEndpointPolicyDefinitions ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ServiceEndpointPolicyDefinitionList.json
     */
    /**
     * Sample code: List service endpoint definitions in service end point policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listServiceEndpointDefinitionsInServiceEndPointPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceEndpointPolicyDefinitions().listByResourceGroup("rg1",
            "testPolicy", com.azure.core.util.Context.NONE);
    }
}
