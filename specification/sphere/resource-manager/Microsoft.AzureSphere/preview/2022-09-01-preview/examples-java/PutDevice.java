/** Samples for Devices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PutDevice.json
     */
    /**
     * Sample code: Devices_CreateOrUpdate.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void devicesCreateOrUpdate(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .devices()
            .define(
                "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
            .withExistingDeviceGroup("MyResourceGroup1", "MyCatalog1", "MyProduct1", "myDeviceGroup1")
            .create();
    }
}
