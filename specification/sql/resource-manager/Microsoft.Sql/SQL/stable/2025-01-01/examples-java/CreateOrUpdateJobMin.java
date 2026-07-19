
import com.azure.resourcemanager.sql.fluent.models.JobInner;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateJobMin.json
     */
    /**
     * Sample code: Create a job with default properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createAJobWithDefaultProperties(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobs().createOrUpdateWithResponse("group1", "server1", "agent1", "job1",
            new JobInner(), com.azure.core.util.Context.NONE);
    }
}
