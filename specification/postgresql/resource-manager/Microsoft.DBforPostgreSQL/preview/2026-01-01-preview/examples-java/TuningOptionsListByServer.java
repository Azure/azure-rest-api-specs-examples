
/**
 * Samples for TuningOptionsOperation ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/TuningOptionsListByServer.json
     */
    /**
     * Sample code: List the tuning options of a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        listTheTuningOptionsOfAServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningOptionsOperations().listByServer("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
