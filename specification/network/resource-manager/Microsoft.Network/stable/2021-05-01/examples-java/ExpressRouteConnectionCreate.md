Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ExpressRouteConnectionInner;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitPeeringId;

/** Samples for ExpressRouteConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteConnectionCreate.json
     */
    /**
     * Sample code: ExpressRouteConnectionCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteConnectionCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteConnections()
            .createOrUpdate(
                "resourceGroupName",
                "gateway-2",
                "connectionName",
                new ExpressRouteConnectionInner()
                    .withId(
                        "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName")
                    .withName("connectionName")
                    .withExpressRouteCircuitPeering(
                        new ExpressRouteCircuitPeeringId()
                            .withId(
                                "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"))
                    .withAuthorizationKey("authorizationKey")
                    .withRoutingWeight(2),
                Context.NONE);
    }
}
```
