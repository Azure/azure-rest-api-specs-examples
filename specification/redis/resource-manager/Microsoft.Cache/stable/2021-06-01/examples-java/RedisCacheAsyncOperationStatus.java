import com.azure.core.util.Context;

/** Samples for AsyncOperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheAsyncOperationStatus.json
     */
    /**
     * Sample code: RedisCacheAsyncOperationStatus.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheAsyncOperationStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getAsyncOperationStatus()
            .getWithResponse("East US", "c7ba2bf5-5939-4d79-b037-2964ccf097da", Context.NONE);
    }
}
