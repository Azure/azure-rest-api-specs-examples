
/**
 * Samples for JobCredentials ListByAgent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobCredentialsByAgent.json
     */
    /**
     * Sample code: List credentials in a job agent.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listCredentialsInAJobAgent(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobCredentials().listByAgent("group1", "server1", "agent1",
            com.azure.core.util.Context.NONE);
    }
}
