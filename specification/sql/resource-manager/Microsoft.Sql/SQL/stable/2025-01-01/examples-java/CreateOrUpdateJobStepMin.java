
import com.azure.resourcemanager.sql.fluent.models.JobStepInner;
import com.azure.resourcemanager.sql.models.JobStepAction;

/**
 * Samples for JobSteps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateJobStepMin.json
     */
    /**
     * Sample code: Create or update a job step with minimal properties specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createOrUpdateAJobStepWithMinimalPropertiesSpecified(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobSteps().createOrUpdateWithResponse("group1", "server1", "agent1", "job1", "step1",
            new JobStepInner().withTargetGroup(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup0")
                .withAction(new JobStepAction().withValue("select 1")),
            com.azure.core.util.Context.NONE);
    }
}
