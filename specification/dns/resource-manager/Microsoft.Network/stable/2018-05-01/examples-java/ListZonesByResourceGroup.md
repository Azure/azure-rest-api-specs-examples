Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Zones ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListZonesByResourceGroup.json
     */
    /**
     * Sample code: List zones by resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listZonesByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getZones().listByResourceGroup("rg1", null, Context.NONE);
    }
}
```
