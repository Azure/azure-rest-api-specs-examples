
import com.azure.resourcemanager.sql.fluent.models.SyncGroupInner;
import com.azure.resourcemanager.sql.models.SyncConflictResolutionPolicy;

/**
 * Samples for SyncGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncGroupUpdate.json
     */
    /**
     * Sample code: Update a sync group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASyncGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncGroups().createOrUpdate("syncgroupcrud-65440",
            "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187",
            new SyncGroupInner().withInterval(-1).withConflictResolutionPolicy(SyncConflictResolutionPolicy.HUB_WIN)
                .withSyncDatabaseId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328")
                .withHubDatabaseUsername("hubUser").withUsePrivateLinkConnection(true),
            com.azure.core.util.Context.NONE);
    }
}
