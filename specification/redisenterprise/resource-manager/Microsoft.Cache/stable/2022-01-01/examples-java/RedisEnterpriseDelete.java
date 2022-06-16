import com.azure.core.util.Context;

/** Samples for RedisEnterprise Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDelete.json
     */
    /**
     * Sample code: RedisEnterpriseDelete.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDelete(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.redisEnterprises().delete("rg1", "cache1", Context.NONE);
    }
}
