/** Samples for Devices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetDevice.json
     */
    /**
     * Sample code: Devices_Get.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void devicesGet(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .devices()
            .getWithResponse(
                "MyResourceGroup1",
                "MyCatalog1",
                "MyProduct1",
                "myDeviceGroup1",
                "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
                com.azure.core.util.Context.NONE);
    }
}
