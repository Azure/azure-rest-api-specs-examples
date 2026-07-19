
/**
 * Samples for SyncGroups ListSyncDatabaseIds.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncGroupGetSyncDatabaseId.json
     */
    /**
     * Sample code: Get a sync database ID.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getASyncDatabaseID(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().listSyncDatabaseIds("westus", com.azure.core.util.Context.NONE);
    }
}
