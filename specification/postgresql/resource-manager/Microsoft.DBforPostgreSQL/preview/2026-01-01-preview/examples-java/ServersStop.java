
/**
 * Samples for Servers Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersStop.json
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
