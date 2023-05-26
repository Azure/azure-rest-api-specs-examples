/** Samples for PatchSchedules ListByRedisResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCachePatchSchedulesList.json
     */
    /**
     * Sample code: RedisCachePatchSchedulesList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCachePatchSchedulesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getPatchSchedules()
            .listByRedisResource("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
