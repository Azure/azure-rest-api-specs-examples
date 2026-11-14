
import com.azure.resourcemanager.iothub.models.ImportDevicesRequest;

/**
 * Samples for IotHubResource ImportDevices.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/iothub_importdevices.json
     */
    /**
     * Sample code: IotHubResource_ImportDevices.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceImportDevices(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().importDevicesWithResponse("myResourceGroup", "testHub",
            new ImportDevicesRequest().withInputBlobContainerUri("testBlob").withOutputBlobContainerUri("testBlob"),
            com.azure.core.util.Context.NONE);
    }
}
