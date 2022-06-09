```java
import com.azure.core.util.Context;

/** Samples for PacketCaptures Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherPacketCaptureDelete.json
     */
    /**
     * Sample code: Delete packet capture.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePacketCapture(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPacketCaptures().delete("rg1", "nw1", "pc1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
