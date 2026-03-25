
/**
 * Samples for IntegrationRuntimes ListOutboundNetworkDependenciesEndpoints.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: IntegrationRuntimes_OutboundNetworkDependenciesEndpoints.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesOutboundNetworkDependenciesEndpoints(
        com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().listOutboundNetworkDependenciesEndpointsWithResponse("exampleResourceGroup",
            "exampleFactoryName", "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
