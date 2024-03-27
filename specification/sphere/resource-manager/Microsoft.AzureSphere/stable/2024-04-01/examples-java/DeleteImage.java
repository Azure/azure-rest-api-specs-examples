
/**
 * Samples for Images Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/DeleteImage.json
     */
    /**
     * Sample code: Images_Delete.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void imagesDelete(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.images().delete("MyResourceGroup1", "MyCatalog1", "00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
