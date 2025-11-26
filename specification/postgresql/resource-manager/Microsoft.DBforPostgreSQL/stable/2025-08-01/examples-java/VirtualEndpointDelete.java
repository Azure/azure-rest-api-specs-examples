
/**
 * Samples for VirtualEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * VirtualEndpointDelete.json
     */
    /**
     * Sample code: Delete a pair of virtual endpoints.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        deleteAPairOfVirtualEndpoints(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.virtualEndpoints().delete("exampleresourcegroup", "exampleserver", "examplebasename",
            com.azure.core.util.Context.NONE);
    }
}
