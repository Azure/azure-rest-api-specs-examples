
/**
 * Samples for JobCredentials Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeleteJobCredential.json
     */
    /**
     * Sample code: Delete a credential.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteACredential(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobCredentials().deleteWithResponse("group1", "server1",
            "agent1", "cred1", com.azure.core.util.Context.NONE);
    }
}
