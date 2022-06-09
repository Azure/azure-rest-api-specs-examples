```java
import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void operationsList(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.operations().list(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iotcentral_1.0.0/sdk/iotcentral/azure-resourcemanager-iotcentral/README.md) on how to add the SDK to your project and authenticate.
