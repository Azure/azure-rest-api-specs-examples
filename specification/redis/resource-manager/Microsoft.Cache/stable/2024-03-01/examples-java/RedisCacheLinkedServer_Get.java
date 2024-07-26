
/**
 * Samples for LinkedServer Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheLinkedServer_Get.json
     */
    /**
     * Sample code: LinkedServer_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void linkedServerGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getLinkedServers().getWithResponse("rg1", "cache1", "cache2",
            com.azure.core.util.Context.NONE);
    }
}
