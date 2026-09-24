
/**
 * Samples for ManagedInstancePrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstancePrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Deletes a private endpoint connection with a given name.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        deletesAPrivateEndpointConnectionWithAGivenName(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstancePrivateEndpointConnections().delete("Default", "test-cl",
            "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
