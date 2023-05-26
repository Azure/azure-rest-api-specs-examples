/** Samples for LinkedServer Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheLinkedServer_Delete.json
     */
    /**
     * Sample code: LinkedServerDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void linkedServerDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getLinkedServers()
            .delete("rg1", "cache1", "cache2", com.azure.core.util.Context.NONE);
    }
}
