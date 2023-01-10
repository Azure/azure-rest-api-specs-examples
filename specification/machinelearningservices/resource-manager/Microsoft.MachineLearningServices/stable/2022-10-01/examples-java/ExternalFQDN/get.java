import com.azure.core.util.Context;

/** Samples for Workspaces ListOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ExternalFQDN/get.json
     */
    /**
     * Sample code: ListOutboundNetworkDependenciesEndpoints.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listOutboundNetworkDependenciesEndpoints(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .workspaces()
            .listOutboundNetworkDependenciesEndpointsWithResponse("workspace-1234", "testworkspace", Context.NONE);
    }
}
