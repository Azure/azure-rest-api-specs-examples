
import com.azure.core.util.Context;

/** Samples for SyncAgents ListLinkedDatabases. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncAgentGetLinkedDatabases.json
     */
    /**
     * Sample code: Get sync agent linked databases.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSyncAgentLinkedDatabases(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncAgents().listLinkedDatabases("syncagentcrud-65440",
            "syncagentcrud-8475", "syncagentcrud-3187", Context.NONE);
    }
}
