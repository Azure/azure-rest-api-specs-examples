/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/PrivateEndpointConnectionsGet.json
     */
    /**
     * Sample code: Get a private endpoint connection.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void getAPrivateEndpointConnection(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse(
                "myResourceGroup",
                "myWorkspace",
                "myWorkspace.23456789-1111-1111-1111-111111111111",
                com.azure.core.util.Context.NONE);
    }
}
