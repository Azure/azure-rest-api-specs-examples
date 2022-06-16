import com.azure.core.util.Context;

/** Samples for Workspaces ListOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ExternalFQDN/get.json
     */
    /**
     * Sample code: ListOutboundNetworkDependenciesEndpoints.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listOutboundNetworkDependenciesEndpoints(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .workspaces()
            .listOutboundNetworkDependenciesEndpointsWithResponse("workspace-1234", "testworkspace", Context.NONE);
    }
}
