
import com.azure.resourcemanager.network.fluent.models.ServiceEndpointPolicyInner;

/**
 * Samples for ServiceEndpointPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceEndpointPolicyCreate.json
     */
    /**
     * Sample code: Create service endpoint policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createServiceEndpointPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceEndpointPolicies().createOrUpdate("rg1", "testPolicy",
            new ServiceEndpointPolicyInner().withLocation("westus"), com.azure.core.util.Context.NONE);
    }
}
