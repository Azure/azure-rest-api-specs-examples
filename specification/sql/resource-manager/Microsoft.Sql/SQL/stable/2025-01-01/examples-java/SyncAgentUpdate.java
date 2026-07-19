
import com.azure.resourcemanager.sql.fluent.models.SyncAgentInner;

/**
 * Samples for SyncAgents CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncAgentUpdate.json
     */
    /**
     * Sample code: Update a sync agent.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateASyncAgent(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncAgents().createOrUpdate("syncagentcrud-65440", "syncagentcrud-8475",
            "syncagentcrud-3187",
            new SyncAgentInner().withSyncDatabaseId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/syncagentcrud-8475/databases/sync"),
            com.azure.core.util.Context.NONE);
    }
}
