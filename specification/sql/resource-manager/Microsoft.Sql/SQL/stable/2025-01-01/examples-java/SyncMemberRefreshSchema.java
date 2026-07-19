
/**
 * Samples for SyncMembers RefreshMemberSchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncMemberRefreshSchema.json
     */
    /**
     * Sample code: Refresh a sync member database schema.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void refreshASyncMemberDatabaseSchema(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncMembers().refreshMemberSchema("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", "syncgroupcrud-4879", com.azure.core.util.Context.NONE);
    }
}
