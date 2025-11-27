
/**
 * Samples for BackupsAutomaticAndOnDemand Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * BackupsAutomaticAndOnDemandDelete.json
     */
    /**
     * Sample code: Delete an on demand backup, given its name.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void deleteAnOnDemandBackupGivenItsName(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backupsAutomaticAndOnDemands().delete("exampleresourcegroup", "exampleserver",
            "ondemandbackup-20250601T183022", com.azure.core.util.Context.NONE);
    }
}
