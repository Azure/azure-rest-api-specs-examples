```java
import com.azure.core.util.Context;

/** Samples for VirtualNetworks CheckIpAddressAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkCheckIPAddressAvailability.json
     */
    /**
     * Sample code: Check IP address availability.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkIPAddressAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworks()
            .checkIpAddressAvailabilityWithResponse("rg1", "test-vnet", "10.0.1.4", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
