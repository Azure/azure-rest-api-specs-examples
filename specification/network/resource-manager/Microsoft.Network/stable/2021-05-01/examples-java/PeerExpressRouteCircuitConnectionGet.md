Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PeerExpressRouteCircuitConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PeerExpressRouteCircuitConnectionGet.json
     */
    /**
     * Sample code: PeerExpressRouteCircuitConnectionGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void peerExpressRouteCircuitConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPeerExpressRouteCircuitConnections()
            .getWithResponse(
                "rg1",
                "ExpressRouteARMCircuitA",
                "AzurePrivatePeering",
                "60aee347-e889-4a42-8c1b-0aae8b1e4013",
                Context.NONE);
    }
}
```
