
import com.azure.resourcemanager.sql.fluent.models.JobAgentInner;

/**
 * Samples for JobAgents CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateJobAgent.json
     */
    /**
     * Sample code: Create or update a job agent.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAJobAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobAgents().createOrUpdate("group1", "server1", "agent1",
            new JobAgentInner().withLocation("southeastasia").withDatabaseId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/databases/db1"),
            com.azure.core.util.Context.NONE);
    }
}
