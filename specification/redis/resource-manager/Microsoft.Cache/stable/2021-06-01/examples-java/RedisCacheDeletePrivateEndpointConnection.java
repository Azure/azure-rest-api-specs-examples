import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: RedisCacheDeletePrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheDeletePrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .deleteWithResponse("rgtest01", "cachetest01", "pectest01", Context.NONE);
    }
}
