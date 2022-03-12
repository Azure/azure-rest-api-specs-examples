Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.EffectiveRoutesParameters;

/** Samples for VirtualHubs GetEffectiveVirtualHubRoutes. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/EffectiveRoutesListForConnection.json
     */
    /**
     * Sample code: Effective Routes for a Connection resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void effectiveRoutesForAConnectionResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualHubs()
            .getEffectiveVirtualHubRoutes(
                "rg1",
                "virtualHub1",
                new EffectiveRoutesParameters()
                    .withResourceId(
                        "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName")
                    .withVirtualWanResourceType("ExpressRouteConnection"),
                Context.NONE);
    }
}
```
