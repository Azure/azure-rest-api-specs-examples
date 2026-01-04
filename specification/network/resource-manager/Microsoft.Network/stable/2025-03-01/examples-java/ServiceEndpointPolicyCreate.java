
import com.azure.resourcemanager.network.fluent.models.ServiceEndpointPolicyInner;

/**
 * Samples for ServiceEndpointPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ServiceEndpointPolicyCreate.
     * json
     */
    /**
     * Sample code: Create service endpoint policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createServiceEndpointPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceEndpointPolicies().createOrUpdate("rg1", "testPolicy",
            new ServiceEndpointPolicyInner().withLocation("westus"), com.azure.core.util.Context.NONE);
    }
}
