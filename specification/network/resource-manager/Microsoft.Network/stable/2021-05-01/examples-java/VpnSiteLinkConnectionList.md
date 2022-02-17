Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VpnLinkConnections ListByVpnConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnSiteLinkConnectionList.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteLinkConnectionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnLinkConnections()
            .listByVpnConnection("rg1", "gateway1", "vpnConnection1", Context.NONE);
    }
}
```
