
/**
 * Samples for JobCredentials Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeleteJobCredential.json
     */
    /**
     * Sample code: Delete a credential.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteACredential(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobCredentials().deleteWithResponse("group1", "server1", "agent1", "cred1",
            com.azure.core.util.Context.NONE);
    }
}
