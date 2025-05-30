
import com.azure.resourcemanager.redis.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.redis.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.redis.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Put.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/
     * RedisCachePutPrivateEndpointConnection.json
     */
    /**
     * Sample code: RedisCachePutPrivateEndpointConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCachePutPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getPrivateEndpointConnections().put("rgtest01", "cachetest01",
            "pectest01",
            new PrivateEndpointConnectionInner().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Auto-Approved")),
            com.azure.core.util.Context.NONE);
    }
}
