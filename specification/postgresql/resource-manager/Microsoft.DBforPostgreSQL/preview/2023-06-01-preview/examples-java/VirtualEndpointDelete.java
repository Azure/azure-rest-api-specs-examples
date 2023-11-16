/** Samples for VirtualEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/VirtualEndpointDelete.json
     */
    /**
     * Sample code: Delete a virtual endpoint.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void deleteAVirtualEndpoint(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .virtualEndpoints()
            .delete("testrg", "pgtestsvc4", "pgVirtualEndpoint1", com.azure.core.util.Context.NONE);
    }
}
