```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.models.ImportDevicesRequest;

/** Samples for IotHubResource ImportDevices. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_importdevices.json
     */
    /**
     * Sample code: IotHubResource_ImportDevices.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceImportDevices(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .importDevicesWithResponse(
                "myResourceGroup",
                "testHub",
                new ImportDevicesRequest().withInputBlobContainerUri("testBlob").withOutputBlobContainerUri("testBlob"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iothub_1.2.0-beta.1/sdk/iothub/azure-resourcemanager-iothub/README.md) on how to add the SDK to your project and authenticate.
