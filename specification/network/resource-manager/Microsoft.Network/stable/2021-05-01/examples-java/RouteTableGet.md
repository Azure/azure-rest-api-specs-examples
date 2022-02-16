Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for RouteTables GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/RouteTableGet.json
     */
    /**
     * Sample code: Get route table.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRouteTable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getRouteTables()
            .getByResourceGroupWithResponse("rg1", "testrt", null, Context.NONE);
    }
}
```
