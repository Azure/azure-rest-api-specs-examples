```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCircuitConnectionInner;
import com.azure.resourcemanager.network.models.Ipv6CircuitConnectionConfig;

/** Samples for ExpressRouteCircuitConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitConnectionCreate.json
     */
    /**
     * Sample code: ExpressRouteCircuitConnectionCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteCircuitConnectionCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuitConnections()
            .createOrUpdate(
                "rg1",
                "ExpressRouteARMCircuitA",
                "AzurePrivatePeering",
                "circuitConnectionUSAUS",
                new ExpressRouteCircuitConnectionInner()
                    .withExpressRouteCircuitPeering(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid1/resourceGroups/dedharcktinit/providers/Microsoft.Network/expressRouteCircuits/dedharcktlocal/peerings/AzurePrivatePeering"))
                    .withPeerExpressRouteCircuitPeering(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid2/resourceGroups/dedharcktpeer/providers/Microsoft.Network/expressRouteCircuits/dedharcktremote/peerings/AzurePrivatePeering"))
                    .withAddressPrefix("10.0.0.0/29")
                    .withAuthorizationKey("946a1918-b7a2-4917-b43c-8c4cdaee006a")
                    .withIpv6CircuitConnectionConfig(
                        new Ipv6CircuitConnectionConfig().withAddressPrefix("aa:bb::/125")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
