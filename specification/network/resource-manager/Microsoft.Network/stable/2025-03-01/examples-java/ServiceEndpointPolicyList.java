
/**
 * Samples for ServiceEndpointPolicies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ServiceEndpointPolicyList.
     * json
     */
    /**
     * Sample code: List resource group service endpoint policies.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listResourceGroupServiceEndpointPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceEndpointPolicies().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
