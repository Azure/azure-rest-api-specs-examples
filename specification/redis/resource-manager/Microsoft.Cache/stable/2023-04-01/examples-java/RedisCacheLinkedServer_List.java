/** Samples for LinkedServer List. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheLinkedServer_List.json
     */
    /**
     * Sample code: LinkedServer_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void linkedServerList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getLinkedServers()
            .list("rg1", "cache1", com.azure.core.util.Context.NONE);
    }
}
