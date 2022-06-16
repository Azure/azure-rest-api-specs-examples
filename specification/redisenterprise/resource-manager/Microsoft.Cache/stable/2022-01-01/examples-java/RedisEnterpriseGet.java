import com.azure.core.util.Context;

/** Samples for RedisEnterprise GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseGet.json
     */
    /**
     * Sample code: RedisEnterpriseGet.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseGet(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.redisEnterprises().getByResourceGroupWithResponse("rg1", "cache1", Context.NONE);
    }
}
