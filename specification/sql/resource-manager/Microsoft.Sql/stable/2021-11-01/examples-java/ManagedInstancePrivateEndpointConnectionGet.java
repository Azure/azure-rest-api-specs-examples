
/**
 * Samples for ManagedInstancePrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstancePrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstancePrivateEndpointConnections().getWithResponse(
            "Default", "test-cl", "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
