
/**
 * Samples for IotHubResource GetKeysForKeyName.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/iothub/resource-manager/Microsoft.Devices/IoTHub/preview/2025-08-01-preview/examples/iothub_getkey.
     * json
     */
    /**
     * Sample code: IotHubResource_GetKeysForKeyName.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetKeysForKeyName(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getKeysForKeyNameWithResponse("myResourceGroup", "testHub", "iothubowner",
            com.azure.core.util.Context.NONE);
    }
}
