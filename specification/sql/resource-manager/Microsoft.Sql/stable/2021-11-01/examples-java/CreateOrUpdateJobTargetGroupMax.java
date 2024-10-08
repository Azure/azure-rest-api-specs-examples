
import com.azure.resourcemanager.sql.fluent.models.JobTargetGroupInner;
import com.azure.resourcemanager.sql.models.JobTarget;
import com.azure.resourcemanager.sql.models.JobTargetGroupMembershipType;
import com.azure.resourcemanager.sql.models.JobTargetType;
import java.util.Arrays;

/**
 * Samples for JobTargetGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateJobTargetGroupMax.json
     */
    /**
     * Sample code: Create or update a target group with all properties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateATargetGroupWithAllProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobTargetGroups().createOrUpdateWithResponse("group1",
            "server1", "agent1", "targetGroup1",
            new JobTargetGroupInner().withMembers(Arrays.asList(
                new JobTarget().withMembershipType(JobTargetGroupMembershipType.EXCLUDE)
                    .withType(JobTargetType.SQL_DATABASE).withServerName("server1").withDatabaseName("database1"),
                new JobTarget().withMembershipType(JobTargetGroupMembershipType.INCLUDE)
                    .withType(JobTargetType.SQL_SERVER).withServerName("server1")
                    .withRefreshCredential("fakeTokenPlaceholder"),
                new JobTarget().withMembershipType(JobTargetGroupMembershipType.INCLUDE)
                    .withType(JobTargetType.SQL_ELASTIC_POOL).withServerName("server2").withElasticPoolName("pool1")
                    .withRefreshCredential("fakeTokenPlaceholder"),
                new JobTarget().withMembershipType(JobTargetGroupMembershipType.INCLUDE)
                    .withType(JobTargetType.SQL_SHARD_MAP).withServerName("server3").withShardMapName("shardMap1")
                    .withRefreshCredential("fakeTokenPlaceholder"))),
            com.azure.core.util.Context.NONE);
    }
}
