Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.RoutingIntentInner;
import com.azure.resourcemanager.network.models.RoutingPolicy;
import java.util.Arrays;

/** Samples for RoutingIntent CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/RoutingIntentPut.json
     */
    /**
     * Sample code: RouteTablePut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeTablePut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getRoutingIntents()
            .createOrUpdate(
                "rg1",
                "virtualHub1",
                "Intent1",
                new RoutingIntentInner()
                    .withRoutingPolicies(
                        Arrays
                            .asList(
                                new RoutingPolicy()
                                    .withName("InternetTraffic")
                                    .withDestinations(Arrays.asList("Internet"))
                                    .withNextHop(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
                                new RoutingPolicy()
                                    .withName("PrivateTrafficPolicy")
                                    .withDestinations(Arrays.asList("PrivateTraffic"))
                                    .withNextHop(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"))),
                Context.NONE);
    }
}
```
