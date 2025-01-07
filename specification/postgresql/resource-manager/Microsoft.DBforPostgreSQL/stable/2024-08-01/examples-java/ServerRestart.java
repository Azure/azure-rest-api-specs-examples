
/**
 * Samples for Servers Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/ServerRestart.json
     */
    /**
     * Sample code: ServerRestart.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverRestart(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().restart("testrg", "testserver", null, com.azure.core.util.Context.NONE);
    }
}
