/** Samples for Catalogs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/DeleteCatalog.json
     */
    /**
     * Sample code: Catalogs_Delete.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsDelete(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.catalogs().delete("MyResourceGroup1", "MyCatalog1", com.azure.core.util.Context.NONE);
    }
}
