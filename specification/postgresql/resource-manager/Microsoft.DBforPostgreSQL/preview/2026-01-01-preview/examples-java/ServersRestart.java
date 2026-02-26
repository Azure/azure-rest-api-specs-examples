
/**
 * Samples for Servers Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersRestart.json
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
