import com.azure.core.util.Context;

/** Samples for Databases ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesListKeys.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesListKeys.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesListKeys(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().listKeysWithResponse("rg1", "cache1", "default", Context.NONE);
    }
}
