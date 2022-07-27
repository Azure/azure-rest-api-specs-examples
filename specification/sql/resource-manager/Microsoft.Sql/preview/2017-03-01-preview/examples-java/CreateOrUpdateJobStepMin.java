import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.JobStepInner;
import com.azure.resourcemanager.sql.models.JobStepAction;

/** Samples for JobSteps CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/CreateOrUpdateJobStepMin.json
     */
    /**
     * Sample code: Create or update a job step with minimal properties specified.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAJobStepWithMinimalPropertiesSpecified(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobSteps()
            .createOrUpdateWithResponse(
                "group1",
                "server1",
                "agent1",
                "job1",
                "step1",
                new JobStepInner()
                    .withTargetGroup(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup0")
                    .withCredential(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0")
                    .withAction(new JobStepAction().withValue("select 1")),
                Context.NONE);
    }
}
