import com.azure.core.util.Context;

/** Samples for SyncAgents ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/SyncAgentListByServer.json
     */
    /**
     * Sample code: Get sync agents under a server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSyncAgentsUnderAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getSyncAgents()
            .listByServer("syncagentcrud-65440", "syncagentcrud-8475", Context.NONE);
    }
}
