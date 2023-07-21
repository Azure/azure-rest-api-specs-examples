/** Samples for DeviceGroups ListByProduct. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetDeviceGroups.json
     */
    /**
     * Sample code: DeviceGroups_ListByProduct.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deviceGroupsListByProduct(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .deviceGroups()
            .listByProduct(
                "MyResourceGroup1",
                "MyCatalog1",
                "MyProduct1",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
