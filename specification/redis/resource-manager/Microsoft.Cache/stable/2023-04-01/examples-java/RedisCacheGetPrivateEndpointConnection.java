/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: RedisCacheGetPrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheGetPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .getWithResponse("rgtest01", "cachetest01", "pectest01", com.azure.core.util.Context.NONE);
    }
}
