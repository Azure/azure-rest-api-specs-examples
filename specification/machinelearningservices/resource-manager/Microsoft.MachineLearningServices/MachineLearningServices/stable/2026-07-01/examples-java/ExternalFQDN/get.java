
/**
 * Samples for Workspaces ListOutboundNetworkDependenciesEndpoints.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ExternalFQDN/get.json
     */
    /**
     * Sample code: ListOutboundNetworkDependenciesEndpoints.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listOutboundNetworkDependenciesEndpoints(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().listOutboundNetworkDependenciesEndpointsWithResponse("workspace-1234", "testworkspace",
            com.azure.core.util.Context.NONE);
    }
}
