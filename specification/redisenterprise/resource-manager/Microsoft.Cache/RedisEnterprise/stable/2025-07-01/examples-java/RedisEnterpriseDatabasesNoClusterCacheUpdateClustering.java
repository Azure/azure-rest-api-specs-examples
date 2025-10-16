
import com.azure.resourcemanager.redisenterprise.models.ClusteringPolicy;
import com.azure.resourcemanager.redisenterprise.models.Database;
import com.azure.resourcemanager.redisenterprise.models.EvictionPolicy;
import com.azure.resourcemanager.redisenterprise.models.Protocol;

/**
 * Samples for Databases Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/RedisEnterprise/stable/2025-07-01/examples/
     * RedisEnterpriseDatabasesNoClusterCacheUpdateClustering.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesUpdate Clustering on No Cluster Cache.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesUpdateClusteringOnNoClusterCache(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        Database resource = manager.databases()
            .getWithResponse("rg1", "cache1", "default", com.azure.core.util.Context.NONE).getValue();
        resource.update().withClientProtocol(Protocol.ENCRYPTED)
            .withClusteringPolicy(ClusteringPolicy.ENTERPRISE_CLUSTER).withEvictionPolicy(EvictionPolicy.NO_EVICTION)
            .apply();
    }
}
