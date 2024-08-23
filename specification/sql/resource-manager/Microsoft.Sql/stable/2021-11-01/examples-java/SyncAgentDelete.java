
/**
 * Samples for SyncAgents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncAgentDelete.json
     */
    /**
     * Sample code: Delete a sync agent.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteASyncAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncAgents().delete("syncagentcrud-65440", "syncagentcrud-8475",
            "syncagentcrud-3187", com.azure.core.util.Context.NONE);
    }
}
