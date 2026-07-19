
/**
 * Samples for SyncGroups CancelSync.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncGroupCancelSync.json
     */
    /**
     * Sample code: Cancel a sync group synchronization.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void cancelASyncGroupSynchronization(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().cancelSyncWithResponse("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
