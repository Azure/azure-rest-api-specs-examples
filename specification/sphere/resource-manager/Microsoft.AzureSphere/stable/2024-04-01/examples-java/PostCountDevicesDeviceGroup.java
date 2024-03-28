
/**
 * Samples for DeviceGroups CountDevices.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/
     * PostCountDevicesDeviceGroup.json
     */
    /**
     * Sample code: DeviceGroups_CountDevices.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deviceGroupsCountDevices(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.deviceGroups().countDevicesWithResponse("MyResourceGroup1", "MyCatalog1", "MyProduct1",
            "MyDeviceGroup1", com.azure.core.util.Context.NONE);
    }
}
