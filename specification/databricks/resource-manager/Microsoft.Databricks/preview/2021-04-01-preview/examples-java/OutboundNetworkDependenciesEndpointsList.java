/** Samples for OutboundNetworkDependenciesEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/OutboundNetworkDependenciesEndpointsList.json
     */
    /**
     * Sample code: List OutboundNetworkDependenciesEndpoints by Workspace.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void listOutboundNetworkDependenciesEndpointsByWorkspace(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .outboundNetworkDependenciesEndpoints()
            .listWithResponse("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
