
/**
 * Samples for VirtualEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/VirtualEndpointDelete.json
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
