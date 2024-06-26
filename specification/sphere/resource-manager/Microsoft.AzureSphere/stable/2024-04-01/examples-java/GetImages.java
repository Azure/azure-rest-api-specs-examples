
/**
 * Samples for Images ListByCatalog.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetImages.json
     */
    /**
     * Sample code: Images_ListByCatalog.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void imagesListByCatalog(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.images().listByCatalog("MyResourceGroup1", "MyCatalog1", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
