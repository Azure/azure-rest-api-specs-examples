/** Samples for Catalogs ListDevices. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostListDevicesByCatalog.json
     */
    /**
     * Sample code: Catalogs_ListDevices.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsListDevices(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .catalogs()
            .listDevices("MyResourceGroup1", "MyCatalog1", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
