/** Samples for OutboundNetworkDependenciesEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/OutboundNetworkDependenciesEndpointsList.json
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
