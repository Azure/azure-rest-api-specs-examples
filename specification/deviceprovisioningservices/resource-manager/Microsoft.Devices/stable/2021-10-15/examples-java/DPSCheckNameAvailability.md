Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceprovisioningservices_1.1.0-beta.1/sdk/deviceprovisioningservices/azure-resourcemanager-deviceprovisioningservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.deviceprovisioningservices.models.OperationInputs;

/** Samples for IotDpsResource CheckProvisioningServiceNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSCheckNameAvailability.json
     */
    /**
     * Sample code: DPSCheckName.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSCheckName(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .iotDpsResources()
            .checkProvisioningServiceNameAvailabilityWithResponse(
                new OperationInputs().withName("test213123"), Context.NONE);
    }
}
```
