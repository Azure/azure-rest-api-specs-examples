
/**
 * Samples for SyncMembers ListBySyncGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncMemberListBySyncGroup.json
     */
    /**
     * Sample code: List sync members under a sync group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listSyncMembersUnderASyncGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncMembers().listBySyncGroup("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
