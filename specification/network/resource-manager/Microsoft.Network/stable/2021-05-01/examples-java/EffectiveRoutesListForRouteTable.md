Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.EffectiveRoutesParameters;

/** Samples for VirtualHubs GetEffectiveVirtualHubRoutes. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/EffectiveRoutesListForRouteTable.json
     */
    /**
     * Sample code: Effective Routes for a Route Table resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void effectiveRoutesForARouteTableResource(com.azure.resourcemanager.AzureResourceManager azure) {
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
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1")
                    .withVirtualWanResourceType("RouteTable"),
                Context.NONE);
    }
}
```
