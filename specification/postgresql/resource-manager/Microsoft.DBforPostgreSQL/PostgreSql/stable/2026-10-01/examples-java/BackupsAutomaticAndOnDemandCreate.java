
/**
 * Samples for BackupsAutomaticAndOnDemand Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/BackupsAutomaticAndOnDemandCreate.json
     */
    /**
     * Sample code: Create an on demand backup of a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        createAnOnDemandBackupOfAServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backupsAutomaticAndOnDemands().create("exampleresourcegroup", "exampleserver",
            "ondemandbackup-20250601T183022", com.azure.core.util.Context.NONE);
    }
}
