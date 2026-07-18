
import com.azure.resourcemanager.sql.fluent.models.JobTargetGroupInner;
import java.util.Arrays;

/**
 * Samples for JobTargetGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateJobTargetGroupMin.json
     */
    /**
     * Sample code: Create or update a target group with minimal properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createOrUpdateATargetGroupWithMinimalProperties(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobTargetGroups().createOrUpdateWithResponse("group1", "server1", "agent1",
            "targetGroup1", new JobTargetGroupInner().withMembers(Arrays.asList()), com.azure.core.util.Context.NONE);
    }
}
