
/**
 * Samples for SyncGroups ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncGroupListByDatabase.json
     */
    /**
     * Sample code: List sync groups under a given database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listSyncGroupsUnderAGivenDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().listByDatabase("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", com.azure.core.util.Context.NONE);
    }
}
