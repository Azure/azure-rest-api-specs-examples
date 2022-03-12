Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AzureFirewalls GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/AzureFirewallGetWithIpGroups.json
     */
    /**
     * Sample code: Get Azure Firewall With IpGroups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureFirewallWithIpGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getAzureFirewalls()
            .getByResourceGroupWithResponse("rg1", "azurefirewall", Context.NONE);
    }
}
```
