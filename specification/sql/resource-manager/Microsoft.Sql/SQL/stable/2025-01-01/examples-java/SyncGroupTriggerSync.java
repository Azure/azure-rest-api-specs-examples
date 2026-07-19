
/**
 * Samples for SyncGroups TriggerSync.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncGroupTriggerSync.json
     */
    /**
     * Sample code: Trigger a sync group synchronization.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void triggerASyncGroupSynchronization(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().triggerSyncWithResponse("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
