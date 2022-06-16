import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseListPrivateLinkResources.json
     */
    /**
     * Sample code: RedisEnterpriseListPrivateLinkResources.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseListPrivateLinkResources(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.privateLinkResources().listByCluster("rg1", "cache1", Context.NONE);
    }
}
