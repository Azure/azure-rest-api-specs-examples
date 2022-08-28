import com.azure.core.util.Context;

/** Samples for ServiceEndpointPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ServiceEndpointPolicyDelete.json
     */
    /**
     * Sample code: Delete service endpoint policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteServiceEndpointPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getServiceEndpointPolicies()
            .delete("rg1", "serviceEndpointPolicy1", Context.NONE);
    }
}
