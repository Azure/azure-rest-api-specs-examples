Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.RouteInner;
import com.azure.resourcemanager.network.fluent.models.RouteTableInner;
import com.azure.resourcemanager.network.models.RouteNextHopType;
import java.util.Arrays;

/** Samples for RouteTables CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/RouteTableCreateWithRoute.json
     */
    /**
     * Sample code: Create route table with route.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createRouteTableWithRoute(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getRouteTables()
            .createOrUpdate(
                "rg1",
                "testrt",
                new RouteTableInner()
                    .withLocation("westus")
                    .withRoutes(
                        Arrays
                            .asList(
                                new RouteInner()
                                    .withName("route1")
                                    .withAddressPrefix("10.0.3.0/24")
                                    .withNextHopType(RouteNextHopType.VIRTUAL_NETWORK_GATEWAY)))
                    .withDisableBgpRoutePropagation(true),
                Context.NONE);
    }
}
```
