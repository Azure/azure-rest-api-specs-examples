
/**
 * Samples for Servers Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersStart.json
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
