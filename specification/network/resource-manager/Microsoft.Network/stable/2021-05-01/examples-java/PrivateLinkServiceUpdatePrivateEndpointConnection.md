Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.network.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateLinkServices UpdatePrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateLinkServiceUpdatePrivateEndpointConnection.json
     */
    /**
     * Sample code: approve or reject private end point connection for a private link service.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void approveOrRejectPrivateEndPointConnectionForAPrivateLinkService(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateLinkServices()
            .updatePrivateEndpointConnectionWithResponse(
                "rg1",
                "testPls",
                "testPlePeConnection",
                new PrivateEndpointConnectionInner()
                    .withName("testPlePeConnection")
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus("Approved")
                            .withDescription("approved it for some reason.")),
                Context.NONE);
    }
}
```
