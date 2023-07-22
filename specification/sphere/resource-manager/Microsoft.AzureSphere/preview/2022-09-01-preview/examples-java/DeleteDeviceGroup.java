/** Samples for DeviceGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/DeleteDeviceGroup.json
     */
    /**
     * Sample code: DeviceGroups_Delete.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deviceGroupsDelete(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .deviceGroups()
            .delete("MyResourceGroup1", "MyCatalog1", "MyProduct1", "MyDeviceGroup1", com.azure.core.util.Context.NONE);
    }
}
