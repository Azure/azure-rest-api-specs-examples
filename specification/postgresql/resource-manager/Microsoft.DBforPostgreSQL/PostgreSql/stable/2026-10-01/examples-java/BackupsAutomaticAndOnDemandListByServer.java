
/**
 * Samples for BackupsAutomaticAndOnDemand ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/BackupsAutomaticAndOnDemandListByServer.json
     */
    /**
     * Sample code: List all available backups of a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        listAllAvailableBackupsOfAServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backupsAutomaticAndOnDemands().listByServer("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
