Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for NetworkInterfaces ListCloudServiceNetworkInterfaces. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/CloudServiceNetworkInterfaceList.json
     */
    /**
     * Sample code: List cloud service network interfaces.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServiceNetworkInterfaces(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaces()
            .listCloudServiceNetworkInterfaces("rg1", "cs1", Context.NONE);
    }
}
```
