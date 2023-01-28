import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheListPrivateEndpointConnections.json
     */
    /**
     * Sample code: RedisCacheListPrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheListPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .list("rgtest01", "cachetest01", Context.NONE);
    }
}
