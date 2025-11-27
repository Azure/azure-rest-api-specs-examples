
/**
 * Samples for Servers Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersStop.json
     */
    /**
     * Sample code: Stop a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void stopAServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().stop("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE);
    }
}
