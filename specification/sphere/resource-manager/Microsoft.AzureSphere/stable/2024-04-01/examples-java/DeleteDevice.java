
/**
 * Samples for Devices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/DeleteDevice.json
     */
    /**
     * Sample code: Devices_Delete.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void devicesDelete(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.devices().delete("MyResourceGroup1", "MyCatalog1", "MyProductName1", "DeviceGroupName1",
            "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
            com.azure.core.util.Context.NONE);
    }
}
