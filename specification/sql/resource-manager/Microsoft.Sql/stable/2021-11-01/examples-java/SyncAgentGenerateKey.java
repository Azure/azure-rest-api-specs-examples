
/**
 * Samples for SyncAgents GenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncAgentGenerateKey.json
     */
    /**
     * Sample code: Generate a sync agent key.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generateASyncAgentKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncAgents().generateKeyWithResponse("syncagentcrud-65440",
            "syncagentcrud-8475", "syncagentcrud-3187", com.azure.core.util.Context.NONE);
    }
}
