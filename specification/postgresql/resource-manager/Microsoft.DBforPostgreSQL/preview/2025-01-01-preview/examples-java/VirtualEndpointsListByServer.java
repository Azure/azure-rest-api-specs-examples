
/**
 * Samples for VirtualEndpoints ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * VirtualEndpointsListByServer.json
     */
    /**
     * Sample code: VirtualEndpointListByServer.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        virtualEndpointListByServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.virtualEndpoints().listByServer("testrg", "pgtestsvc4", com.azure.core.util.Context.NONE);
    }
}
