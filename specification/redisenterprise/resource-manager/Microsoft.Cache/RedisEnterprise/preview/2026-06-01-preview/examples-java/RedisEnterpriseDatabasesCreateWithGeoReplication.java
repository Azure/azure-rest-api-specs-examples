
import com.azure.resourcemanager.redisenterprise.models.AccessKeysAuthentication;
import com.azure.resourcemanager.redisenterprise.models.ClusteringPolicy;
import com.azure.resourcemanager.redisenterprise.models.DatabasePropertiesGeoReplication;
import com.azure.resourcemanager.redisenterprise.models.EvictionPolicy;
import com.azure.resourcemanager.redisenterprise.models.LinkedDatabase;
import com.azure.resourcemanager.redisenterprise.models.Protocol;
import java.util.Arrays;

/**
 * Samples for Databases Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RedisEnterpriseDatabasesCreateWithGeoReplication.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesCreate With Active Geo Replication.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesCreateWithActiveGeoReplication(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().define("default").withExistingRedisEnterprise("rg1", "cache1")
            .withAccessKeysAuthentication(AccessKeysAuthentication.ENABLED).withClientProtocol(Protocol.ENCRYPTED)
            .withPort(10000).withClusteringPolicy(ClusteringPolicy.ENTERPRISE_CLUSTER)
            .withEvictionPolicy(EvictionPolicy.NO_EVICTION)
            .withGeoReplication(new DatabasePropertiesGeoReplication().withGroupNickname("groupName")
                .withLinkedDatabases(Arrays.asList(new LinkedDatabase().withId(
                    "/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
                    new LinkedDatabase().withId(
                        "/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8e/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default"))))
            .withNotifyKeyspaceEvents("KEA").create();
    }
}
