
/**
 * Samples for Servers Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersRestart.
     * json
     */
    /**
     * Sample code: Restart PostgreSQL database engine in a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void restartPostgreSQLDatabaseEngineInAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().restart("exampleresourcegroup", "exampleserver", null, com.azure.core.util.Context.NONE);
    }
}
