Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for NetworkInterfaces ListEffectiveNetworkSecurityGroups. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkInterfaceEffectiveNSGList.json
     */
    /**
     * Sample code: List network interface effective network security groups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNetworkInterfaceEffectiveNetworkSecurityGroups(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaces()
            .listEffectiveNetworkSecurityGroups("rg1", "nic1", Context.NONE);
    }
}
```
