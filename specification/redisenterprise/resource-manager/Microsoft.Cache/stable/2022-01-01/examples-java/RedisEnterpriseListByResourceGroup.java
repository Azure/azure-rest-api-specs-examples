import com.azure.core.util.Context;

/** Samples for RedisEnterprise ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseListByResourceGroup.json
     */
    /**
     * Sample code: RedisEnterpriseListByResourceGroup.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseListByResourceGroup(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.redisEnterprises().listByResourceGroup("rg1", Context.NONE);
    }
}
