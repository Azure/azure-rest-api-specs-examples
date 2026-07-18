
/**
 * Samples for ManagedInstancePrivateEndpointConnections ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstancePrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsListOfPrivateEndpointConnectionsOnAServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstancePrivateEndpointConnections().listByManagedInstance("Default",
            "test-cl", com.azure.core.util.Context.NONE);
    }
}
