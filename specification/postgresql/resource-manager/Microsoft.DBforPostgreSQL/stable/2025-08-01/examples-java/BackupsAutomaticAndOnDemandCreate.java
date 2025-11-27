
/**
 * Samples for BackupsAutomaticAndOnDemand Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * BackupsAutomaticAndOnDemandCreate.json
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
