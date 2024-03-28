
/**
 * Samples for Devices ListByDeviceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetDevices.json
     */
    /**
     * Sample code: Devices_ListByDeviceGroup.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void devicesListByDeviceGroup(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.devices().listByDeviceGroup("MyResourceGroup1", "MyCatalog1", "MyProduct1", "myDeviceGroup1",
            com.azure.core.util.Context.NONE);
    }
}
