import com.azure.core.util.Context;

/** Samples for ServiceEndpointPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ServiceEndpointPolicyListAll.json
     */
    /**
     * Sample code: List all service endpoint policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllServiceEndpointPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceEndpointPolicies().list(Context.NONE);
    }
}
