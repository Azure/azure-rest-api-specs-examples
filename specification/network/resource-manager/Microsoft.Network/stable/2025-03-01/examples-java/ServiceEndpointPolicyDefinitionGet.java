
/**
 * Samples for ServiceEndpointPolicyDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ServiceEndpointPolicyDefinitionGet.json
     */
    /**
     * Sample code: Get service endpoint definition in service endpoint policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getServiceEndpointDefinitionInServiceEndpointPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceEndpointPolicyDefinitions().getWithResponse("rg1",
            "testPolicy", "testDefinition", com.azure.core.util.Context.NONE);
    }
}
