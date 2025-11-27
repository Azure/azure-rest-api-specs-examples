
/**
 * Samples for BackupsAutomaticAndOnDemand Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * BackupsAutomaticAndOnDemandGet.json
     */
    /**
     * Sample code: Get an on demand backup, given its name.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        getAnOnDemandBackupGivenItsName(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backupsAutomaticAndOnDemands().getWithResponse("exampleresourcegroup", "exampleserver",
            "backup_638830782181266873", com.azure.core.util.Context.NONE);
    }
}
