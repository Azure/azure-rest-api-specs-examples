
import com.azure.resourcemanager.sql.fluent.models.JobStepInner;
import com.azure.resourcemanager.sql.models.JobStepAction;
import com.azure.resourcemanager.sql.models.JobStepActionSource;
import com.azure.resourcemanager.sql.models.JobStepActionType;
import com.azure.resourcemanager.sql.models.JobStepExecutionOptions;
import com.azure.resourcemanager.sql.models.JobStepOutput;
import com.azure.resourcemanager.sql.models.JobStepOutputType;
import java.util.UUID;

/**
 * Samples for JobSteps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateJobStepMax.json
     */
    /**
     * Sample code: Create or update a job step with all properties specified.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateAJobStepWithAllPropertiesSpecified(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobSteps().createOrUpdateWithResponse("group1", "server1",
            "agent1", "job1", "step1",
            new JobStepInner().withStepId(1).withTargetGroup(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup1")
                .withCredential("fakeTokenPlaceholder")
                .withAction(new JobStepAction().withType(JobStepActionType.TSQL).withSource(JobStepActionSource.INLINE)
                    .withValue("select 2"))
                .withOutput(new JobStepOutput().withType(JobStepOutputType.SQL_DATABASE)
                    .withSubscriptionId(UUID.fromString("3501b905-a848-4b5d-96e8-b253f62d735a"))
                    .withResourceGroupName("group3").withServerName("server3").withDatabaseName("database3")
                    .withSchemaName("myschema1234").withTableName("mytable5678").withCredential("fakeTokenPlaceholder"))
                .withExecutionOptions(new JobStepExecutionOptions().withTimeoutSeconds(1234).withRetryAttempts(42)
                    .withInitialRetryIntervalSeconds(11).withMaximumRetryIntervalSeconds(222)
                    .withRetryIntervalBackoffMultiplier(3.0F)),
            com.azure.core.util.Context.NONE);
    }
}
