Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iotcentral_1.1.0-beta.1/sdk/iotcentral/azure-resourcemanager-iotcentral/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.iotcentral.models.AppPatch;
import com.azure.resourcemanager.iotcentral.models.SystemAssignedServiceIdentity;
import com.azure.resourcemanager.iotcentral.models.SystemAssignedServiceIdentityType;

/** Samples for Apps Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Update.json
     */
    /**
     * Sample code: Apps_Update.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsUpdate(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager
            .apps()
            .update(
                "resRg",
                "myIoTCentralApp",
                new AppPatch()
                    .withIdentity(
                        new SystemAssignedServiceIdentity().withType(SystemAssignedServiceIdentityType.SYSTEM_ASSIGNED))
                    .withDisplayName("My IoT Central App 2"),
                Context.NONE);
    }
}
```
