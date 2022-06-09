```java
import com.azure.resourcemanager.deviceupdate.models.AuthenticationType;
import com.azure.resourcemanager.deviceupdate.models.DiagnosticStorageProperties;
import com.azure.resourcemanager.deviceupdate.models.IotHubSettings;
import java.util.Arrays;

/** Samples for Instances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Instances/Instances_Create.json
     */
    /**
     * Sample code: Creates or updates Instance.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void createsOrUpdatesInstance(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager
            .instances()
            .define("blue")
            .withRegion("westus2")
            .withExistingAccount("test-rg", "contoso")
            .withIotHubs(
                Arrays
                    .asList(
                        new IotHubSettings()
                            .withResourceId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Devices/IotHubs/blue-contoso-hub")))
            .withEnableDiagnostics(false)
            .withDiagnosticStorageProperties(
                new DiagnosticStorageProperties()
                    .withAuthenticationType(AuthenticationType.KEY_BASED)
                    .withConnectionString("string")
                    .withResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/adu-resource-group/providers/Microsoft.Storage/storageAccounts/testAccount"))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceupdate_1.0.0-beta.1/sdk/deviceupdate/azure-resourcemanager-deviceupdate/README.md) on how to add the SDK to your project and authenticate.
