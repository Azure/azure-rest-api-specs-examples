
/**
 * Samples for BackupsAutomaticAndOnDemand Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/BackupsAutomaticAndOnDemandDelete.json
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
