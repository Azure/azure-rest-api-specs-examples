import com.azure.resourcemanager.redisenterprise.models.AofFrequency;
import com.azure.resourcemanager.redisenterprise.models.ClusteringPolicy;
import com.azure.resourcemanager.redisenterprise.models.EvictionPolicy;
import com.azure.resourcemanager.redisenterprise.models.Module;
import com.azure.resourcemanager.redisenterprise.models.Persistence;
import com.azure.resourcemanager.redisenterprise.models.Protocol;
import java.util.Arrays;

/** Samples for Databases Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseDatabasesCreate.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesCreate.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesCreate(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager
            .databases()
            .define("default")
            .withExistingRedisEnterprise("rg1", "cache1")
            .withClientProtocol(Protocol.ENCRYPTED)
            .withPort(10000)
            .withClusteringPolicy(ClusteringPolicy.ENTERPRISE_CLUSTER)
            .withEvictionPolicy(EvictionPolicy.ALL_KEYS_LRU)
            .withPersistence(new Persistence().withAofEnabled(true).withAofFrequency(AofFrequency.ONES))
            .withModules(
                Arrays
                    .asList(
                        new Module().withName("RedisBloom").withArgs("ERROR_RATE 0.00 INITIAL_SIZE 400"),
                        new Module().withName("RedisTimeSeries").withArgs("RETENTION_POLICY 20"),
                        new Module().withName("RediSearch")))
            .create();
    }
}
