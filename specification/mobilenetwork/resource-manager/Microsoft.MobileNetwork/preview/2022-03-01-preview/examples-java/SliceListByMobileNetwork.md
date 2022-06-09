```java
import com.azure.core.util.Context;

/** Samples for Slices ListByMobileNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SliceListByMobileNetwork.json
     */
    /**
     * Sample code: List mobile network slices in a mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listMobileNetworkSlicesInAMobileNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.slices().listByMobileNetwork("rg1", "testMobileNetwork", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.2/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.
