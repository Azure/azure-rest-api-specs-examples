
/**
 * Samples for SyncMembers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncMemberDelete.json
     */
    /**
     * Sample code: Delete a sync member.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteASyncMember(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncMembers().delete("syncgroupcrud-65440",
            "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", "syncgroupcrud-4879",
            com.azure.core.util.Context.NONE);
    }
}
