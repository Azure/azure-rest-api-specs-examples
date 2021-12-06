Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iotcentral_1.0.0/sdk/iotcentral/azure-resourcemanager-iotcentral/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.iotcentral.models.OperationInputs;

/** Samples for Apps CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_CheckNameAvailability.json
     */
    /**
     * Sample code: Apps_CheckNameAvailability.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsCheckNameAvailability(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager
            .apps()
            .checkNameAvailabilityWithResponse(
                new OperationInputs().withName("myiotcentralapp").withType("IoTApps"), Context.NONE);
    }
}
```
