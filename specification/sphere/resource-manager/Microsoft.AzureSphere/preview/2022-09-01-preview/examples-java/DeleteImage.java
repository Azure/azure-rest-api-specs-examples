/** Samples for Images Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/DeleteImage.json
     */
    /**
     * Sample code: Images_Delete.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void imagesDelete(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.images().delete("MyResourceGroup1", "MyCatalog1", "imageID", com.azure.core.util.Context.NONE);
    }
}
