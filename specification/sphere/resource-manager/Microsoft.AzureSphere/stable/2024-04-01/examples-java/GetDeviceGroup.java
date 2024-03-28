
/**
 * Samples for DeviceGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetDeviceGroup.json
     */
    /**
     * Sample code: DeviceGroups_Get.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deviceGroupsGet(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.deviceGroups().getWithResponse("MyResourceGroup1", "MyCatalog1", "MyProduct1", "MyDeviceGroup1",
            com.azure.core.util.Context.NONE);
    }
}
