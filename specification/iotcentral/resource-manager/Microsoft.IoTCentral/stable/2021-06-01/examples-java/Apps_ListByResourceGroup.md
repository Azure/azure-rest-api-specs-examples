Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iotcentral_1.0.0/sdk/iotcentral/azure-resourcemanager-iotcentral/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Apps ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_ListByResourceGroup.json
     */
    /**
     * Sample code: Apps_ListByResourceGroup.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsListByResourceGroup(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.apps().listByResourceGroup("resRg", Context.NONE);
    }
}
```
