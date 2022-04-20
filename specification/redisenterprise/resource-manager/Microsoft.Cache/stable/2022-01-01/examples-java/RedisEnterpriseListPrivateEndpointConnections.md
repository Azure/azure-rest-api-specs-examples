Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-redisenterprise_1.1.0-beta.1/sdk/redisenterprise/azure-resourcemanager-redisenterprise/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseListPrivateEndpointConnections.json
     */
    /**
     * Sample code: RedisEnterpriseListPrivateEndpointConnections.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseListPrivateEndpointConnections(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.privateEndpointConnections().list("rg1", "cache1", Context.NONE);
    }
}
```
