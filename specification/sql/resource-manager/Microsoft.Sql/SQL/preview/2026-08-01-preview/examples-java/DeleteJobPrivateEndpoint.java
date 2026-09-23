
/**
 * Samples for JobPrivateEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DeleteJobPrivateEndpoint.json
     */
    /**
     * Sample code: Delete a private endpoint.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAPrivateEndpoint(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobPrivateEndpoints().delete("group1", "server1", "agent1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
