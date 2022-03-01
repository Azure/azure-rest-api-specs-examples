Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.1/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AttachedDataNetworks ListByPacketCoreDataPlane. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/AttachedDataNetworkListByPacketCoreDataPlane.json
     */
    /**
     * Sample code: List attached data networks in a data plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listAttachedDataNetworksInADataPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .attachedDataNetworks()
            .listByPacketCoreDataPlane("rg1", "TestPacketCoreCP", "TestPacketCoreDP", Context.NONE);
    }
}
```
