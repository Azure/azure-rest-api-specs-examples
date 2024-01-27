
import com.azure.core.util.Context;

/** Samples for SyncAgents Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncAgentGet.json
     */
    /**
     * Sample code: Get a sync agent.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASyncAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncAgents().getWithResponse("syncagentcrud-65440",
            "syncagentcrud-8475", "syncagentcrud-3187", Context.NONE);
    }
}
