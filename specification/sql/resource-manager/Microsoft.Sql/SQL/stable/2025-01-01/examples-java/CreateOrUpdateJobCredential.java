
import com.azure.resourcemanager.sql.fluent.models.JobCredentialInner;

/**
 * Samples for JobCredentials CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateJobCredential.json
     */
    /**
     * Sample code: Create or update a credential.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateACredential(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobCredentials().createOrUpdateWithResponse("group1", "server1", "agent1", "cred1",
            new JobCredentialInner().withUsername("myuser").withPassword("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
