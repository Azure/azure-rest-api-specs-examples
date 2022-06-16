import com.azure.core.util.Context;

/** Samples for IntegrationRuntimes ListOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: Get outbound network dependency endpoints.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getOutboundNetworkDependencyEndpoints(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimes()
            .listOutboundNetworkDependenciesEndpointsWithResponse(
                "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", Context.NONE);
    }
}
