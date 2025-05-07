
import com.azure.resourcemanager.redisenterprise.models.ClusteringPolicy;
import com.azure.resourcemanager.redisenterprise.models.EvictionPolicy;
import com.azure.resourcemanager.redisenterprise.models.Protocol;

/**
 * Samples for Databases Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/
     * RedisEnterpriseDatabasesNoClusterCacheCreate.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesCreate No Cluster Cache.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesCreateNoClusterCache(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().define("default").withExistingRedisEnterprise("rg1", "cache1")
            .withClientProtocol(Protocol.ENCRYPTED).withPort(10000).withClusteringPolicy(ClusteringPolicy.NO_CLUSTER)
            .withEvictionPolicy(EvictionPolicy.NO_EVICTION).create();
    }
}
