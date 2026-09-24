
/**
 * Samples for SyncGroups RefreshHubSchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/SyncGroupRefreshHubSchema.json
     */
    /**
     * Sample code: Refresh a hub database schema.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void refreshAHubDatabaseSchema(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().refreshHubSchema("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
