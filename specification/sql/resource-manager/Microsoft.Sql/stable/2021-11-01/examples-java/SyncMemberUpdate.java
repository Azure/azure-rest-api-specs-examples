
import com.azure.resourcemanager.sql.fluent.models.SyncMemberInner;
import com.azure.resourcemanager.sql.models.SyncDirection;
import com.azure.resourcemanager.sql.models.SyncMemberDbType;

/**
 * Samples for SyncMembers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncMemberUpdate.json
     */
    /**
     * Sample code: Update a sync member.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASyncMember(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncMembers().createOrUpdate("syncgroupcrud-65440",
            "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", "syncmembercrud-4879",
            new SyncMemberInner().withDatabaseType(SyncMemberDbType.AZURE_SQL_DATABASE)
                .withSyncMemberAzureDatabaseResourceId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328")
                .withUsePrivateLinkConnection(true).withServerName("syncgroupcrud-3379.database.windows.net")
                .withDatabaseName("syncgroupcrud-7421").withUsername("myUser")
                .withSyncDirection(SyncDirection.BIDIRECTIONAL),
            com.azure.core.util.Context.NONE);
    }
}
