import com.azure.resourcemanager.iothub.models.ExportDevicesRequest;

/** Samples for IotHubResource ExportDevices. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_exportdevices.json
     */
    /**
     * Sample code: IotHubResource_ExportDevices.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceExportDevices(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .exportDevicesWithResponse(
                "myResourceGroup",
                "testHub",
                new ExportDevicesRequest().withExportBlobContainerUri("testBlob").withExcludeKeys(true),
                com.azure.core.util.Context.NONE);
    }
}
