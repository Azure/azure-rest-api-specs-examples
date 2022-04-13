Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkServices DeletePrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateLinkServiceDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: delete private end point connection for a private link service.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePrivateEndPointConnectionForAPrivateLinkService(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateLinkServices()
            .deletePrivateEndpointConnection("rg1", "testPls", "testPlePeConnection", Context.NONE);
    }
}
```
