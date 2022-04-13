Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iotcentral_1.1.0-beta.1/sdk/iotcentral/azure-resourcemanager-iotcentral/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateLinks_Get.json
     */
    /**
     * Sample code: PrivateLinks_Get.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void privateLinksGet(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.privateLinks().getWithResponse("resRg", "myIoTCentralApp", "iotApp", Context.NONE);
    }
}
```
