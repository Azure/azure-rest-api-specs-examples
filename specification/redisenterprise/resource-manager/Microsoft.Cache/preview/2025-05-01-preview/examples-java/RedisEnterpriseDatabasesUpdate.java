
import com.azure.resourcemanager.redisenterprise.models.AccessKeysAuthentication;
import com.azure.resourcemanager.redisenterprise.models.Database;
import com.azure.resourcemanager.redisenterprise.models.EvictionPolicy;
import com.azure.resourcemanager.redisenterprise.models.Persistence;
import com.azure.resourcemanager.redisenterprise.models.Protocol;
import com.azure.resourcemanager.redisenterprise.models.RdbFrequency;

/**
 * Samples for Databases Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/
     * RedisEnterpriseDatabasesUpdate.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesUpdate.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        redisEnterpriseDatabasesUpdate(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        Database resource = manager.databases()
            .getWithResponse("rg1", "cache1", "default", com.azure.core.util.Context.NONE).getValue();
        resource.update().withClientProtocol(Protocol.ENCRYPTED).withEvictionPolicy(EvictionPolicy.ALL_KEYS_LRU)
            .withPersistence(new Persistence().withRdbEnabled(true).withRdbFrequency(RdbFrequency.ONE_TWOH))
            .withAccessKeysAuthentication(AccessKeysAuthentication.ENABLED).apply();
    }
}
