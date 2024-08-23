
/**
 * Samples for SyncGroups TriggerSync.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncGroupTriggerSync.json
     */
    /**
     * Sample code: Trigger a sync group synchronization.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void triggerASyncGroupSynchronization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncGroups().triggerSyncWithResponse("syncgroupcrud-65440",
            "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
