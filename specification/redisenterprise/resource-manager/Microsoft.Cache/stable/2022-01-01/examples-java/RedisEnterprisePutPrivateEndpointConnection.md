Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-redisenterprise_1.1.0-beta.1/sdk/redisenterprise/azure-resourcemanager-redisenterprise/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.redisenterprise.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.redisenterprise.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterprisePutPrivateEndpointConnection.json
     */
    /**
     * Sample code: RedisEnterprisePutPrivateEndpointConnection.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterprisePutPrivateEndpointConnection(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager
            .privateEndpointConnections()
            .define("pectest01")
            .withExistingRedisEnterprise("rg1", "cache1")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Auto-Approved"))
            .create();
    }
}
```
