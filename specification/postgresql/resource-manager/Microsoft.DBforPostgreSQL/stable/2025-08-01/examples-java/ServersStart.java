
/**
 * Samples for Servers Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersStart.json
     */
    /**
     * Sample code: Start a stopped server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        startAStoppedServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().start("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE);
    }
}
