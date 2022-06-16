import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.models.ExportDevicesRequest;

/** Samples for IotHubResource ExportDevices. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_exportdevices.json
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
                Context.NONE);
    }
}
