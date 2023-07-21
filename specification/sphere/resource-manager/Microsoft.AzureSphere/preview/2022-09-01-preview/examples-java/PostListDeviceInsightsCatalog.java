/** Samples for Catalogs ListDeviceInsights. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostListDeviceInsightsCatalog.json
     */
    /**
     * Sample code: Catalogs_ListDeviceInsights.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsListDeviceInsights(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .catalogs()
            .listDeviceInsights(
                "MyResourceGroup1", "MyCatalog1", null, 10, null, null, com.azure.core.util.Context.NONE);
    }
}
