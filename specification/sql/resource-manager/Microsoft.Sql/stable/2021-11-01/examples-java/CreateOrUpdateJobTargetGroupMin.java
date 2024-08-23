
import com.azure.resourcemanager.sql.fluent.models.JobTargetGroupInner;
import java.util.Arrays;

/**
 * Samples for JobTargetGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateJobTargetGroupMin.json
     */
    /**
     * Sample code: Create or update a target group with minimal properties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateATargetGroupWithMinimalProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobTargetGroups().createOrUpdateWithResponse("group1",
            "server1", "agent1", "targetGroup1", new JobTargetGroupInner().withMembers(Arrays.asList()),
            com.azure.core.util.Context.NONE);
    }
}
