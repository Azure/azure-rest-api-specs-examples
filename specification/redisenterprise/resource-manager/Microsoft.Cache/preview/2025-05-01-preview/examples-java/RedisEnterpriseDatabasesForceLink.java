
import com.azure.resourcemanager.redisenterprise.models.ForceLinkParameters;
import com.azure.resourcemanager.redisenterprise.models.ForceLinkParametersGeoReplication;
import com.azure.resourcemanager.redisenterprise.models.LinkedDatabase;
import java.util.Arrays;

/**
 * Samples for Databases ForceLinkToReplicationGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/
     * RedisEnterpriseDatabasesForceLink.json
     */
    /**
     * Sample code: How to relink a database after a regional outage.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void howToRelinkADatabaseAfterARegionalOutage(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().forceLinkToReplicationGroup("rg1", "cache1", "default",
            new ForceLinkParameters().withGeoReplication(new ForceLinkParametersGeoReplication()
                .withGroupNickname("groupName")
                .withLinkedDatabases(Arrays.asList(new LinkedDatabase().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
                    new LinkedDatabase().withId(
                        "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default")))),
            com.azure.core.util.Context.NONE);
    }
}
