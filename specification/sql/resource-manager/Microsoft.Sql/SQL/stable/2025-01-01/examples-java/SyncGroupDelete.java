
/**
 * Samples for SyncGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncGroupDelete.json
     */
    /**
     * Sample code: Delete a sync group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteASyncGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().delete("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
