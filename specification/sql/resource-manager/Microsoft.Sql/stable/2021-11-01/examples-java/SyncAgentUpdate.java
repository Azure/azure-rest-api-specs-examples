
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.SyncAgentInner;

/** Samples for SyncAgents CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncAgentUpdate.json
     */
    /**
     * Sample code: Update a sync agent.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASyncAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncAgents().createOrUpdate("syncagentcrud-65440",
            "syncagentcrud-8475", "syncagentcrud-3187",
            new SyncAgentInner().withSyncDatabaseId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/syncagentcrud-8475/databases/sync"),
            Context.NONE);
    }
}
