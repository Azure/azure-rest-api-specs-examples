
/**
 * Samples for ManagedInstancePrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstancePrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Deletes a private endpoint connection with a given name.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deletesAPrivateEndpointConnectionWithAGivenName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstancePrivateEndpointConnections().delete("Default",
            "test-cl", "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
