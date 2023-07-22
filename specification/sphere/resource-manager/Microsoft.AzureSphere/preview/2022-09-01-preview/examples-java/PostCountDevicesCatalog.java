/** Samples for Catalogs CountDevices. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostCountDevicesCatalog.json
     */
    /**
     * Sample code: Catalogs_CountDevices.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsCountDevices(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.catalogs().countDevicesWithResponse("MyResourceGroup1", "MyCatalog1", com.azure.core.util.Context.NONE);
    }
}
