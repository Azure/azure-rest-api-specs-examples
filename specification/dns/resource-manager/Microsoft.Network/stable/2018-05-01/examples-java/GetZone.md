Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Zones GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetZone.json
     */
    /**
     * Sample code: Get zone.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getZones()
            .getByResourceGroupWithResponse("rg1", "zone1", Context.NONE);
    }
}
```
