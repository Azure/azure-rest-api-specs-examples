
import com.azure.core.util.Context;

/** Samples for SyncMembers ListBySyncGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncMemberListBySyncGroup.json
     */
    /**
     * Sample code: List sync members under a sync group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSyncMembersUnderASyncGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncMembers().listBySyncGroup("syncgroupcrud-65440",
            "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", Context.NONE);
    }
}
