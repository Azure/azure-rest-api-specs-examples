import com.azure.core.util.Context;

/** Samples for IotHubResource GetKeysForKeyName. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/examples/iothub_getkey.json
     */
    /**
     * Sample code: IotHubResource_GetKeysForKeyName.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetKeysForKeyName(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .getKeysForKeyNameWithResponse("myResourceGroup", "testHub", "iothubowner", Context.NONE);
    }
}
