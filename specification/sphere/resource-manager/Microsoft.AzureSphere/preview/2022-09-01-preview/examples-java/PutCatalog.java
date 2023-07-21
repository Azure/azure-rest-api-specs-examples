/** Samples for Catalogs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PutCatalog.json
     */
    /**
     * Sample code: Catalogs_CreateOrUpdate.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsCreateOrUpdate(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .catalogs()
            .define("MyCatalog1")
            .withRegion("global")
            .withExistingResourceGroup("MyResourceGroup1")
            .create();
    }
}
