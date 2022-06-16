import com.azure.core.util.Context;

/** Samples for RedisEnterprise List. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseList.json
     */
    /**
     * Sample code: RedisEnterpriseList.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseList(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.redisEnterprises().list(Context.NONE);
    }
}
