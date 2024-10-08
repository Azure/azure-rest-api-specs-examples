
/**
 * Samples for ManagedInstancePrivateEndpointConnections ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstancePrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsListOfPrivateEndpointConnectionsOnAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstancePrivateEndpointConnections()
            .listByManagedInstance("Default", "test-cl", com.azure.core.util.Context.NONE);
    }
}
