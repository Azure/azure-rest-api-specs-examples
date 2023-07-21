/** Samples for Products GenerateDefaultDeviceGroups. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostGenerateDefaultDeviceGroups.json
     */
    /**
     * Sample code: Products_GenerateDefaultDeviceGroups.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void productsGenerateDefaultDeviceGroups(
        com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .products()
            .generateDefaultDeviceGroups(
                "MyResourceGroup1", "MyCatalog1", "MyProduct1", com.azure.core.util.Context.NONE);
    }
}
