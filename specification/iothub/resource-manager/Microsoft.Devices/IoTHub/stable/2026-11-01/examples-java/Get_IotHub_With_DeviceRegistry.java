
/**
 * Samples for IotHubResource GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Get_IotHub_With_DeviceRegistry.json
     */
    /**
     * Sample code: Get_IotHub_With_DeviceRegistry.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void getIotHubWithDeviceRegistry(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getByResourceGroupWithResponse("myResourceGroup", "testHub",
            com.azure.core.util.Context.NONE);
    }
}
