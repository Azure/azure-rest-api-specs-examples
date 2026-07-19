
/**
 * Samples for SyncMembers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncMemberDelete.json
     */
    /**
     * Sample code: Delete a sync member.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteASyncMember(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncMembers().delete("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", "syncgroupcrud-4879", com.azure.core.util.Context.NONE);
    }
}
