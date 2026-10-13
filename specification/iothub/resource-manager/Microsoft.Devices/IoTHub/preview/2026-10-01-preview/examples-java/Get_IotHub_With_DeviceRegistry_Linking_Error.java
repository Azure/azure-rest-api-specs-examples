
/**
 * Samples for IotHubResource GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Get_IotHub_With_DeviceRegistry_Linking_Error.json
     */
    /**
     * Sample code: Get_IotHub_With_DeviceRegistry_Linking_Error.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void getIotHubWithDeviceRegistryLinkingError(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getByResourceGroupWithResponse("myResourceGroup", "testHub",
            com.azure.core.util.Context.NONE);
    }
}
