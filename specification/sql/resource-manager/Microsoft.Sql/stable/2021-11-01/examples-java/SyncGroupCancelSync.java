
/**
 * Samples for SyncGroups CancelSync.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncGroupCancelSync.json
     */
    /**
     * Sample code: Cancel a sync group synchronization.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cancelASyncGroupSynchronization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncGroups().cancelSyncWithResponse("syncgroupcrud-65440",
            "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
