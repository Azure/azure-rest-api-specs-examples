
import com.azure.resourcemanager.redisenterprise.models.AccessKeyType;
import com.azure.resourcemanager.redisenterprise.models.RegenerateKeyParameters;

/**
 * Samples for Databases RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/
     * RedisEnterpriseDatabasesRegenerateKey.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesRegenerateKey.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesRegenerateKey(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().regenerateKey("rg1", "cache1", "default",
            new RegenerateKeyParameters().withKeyType(AccessKeyType.PRIMARY), com.azure.core.util.Context.NONE);
    }
}
