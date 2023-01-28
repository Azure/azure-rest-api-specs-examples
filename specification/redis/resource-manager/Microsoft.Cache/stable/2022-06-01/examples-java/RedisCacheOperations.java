import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheOperations.json
     */
    /**
     * Sample code: RedisCacheOperations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getOperations().list(Context.NONE);
    }
}
