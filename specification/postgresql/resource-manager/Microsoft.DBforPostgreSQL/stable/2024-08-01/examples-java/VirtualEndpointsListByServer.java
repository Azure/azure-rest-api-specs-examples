
/**
 * Samples for VirtualEndpoints ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
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
