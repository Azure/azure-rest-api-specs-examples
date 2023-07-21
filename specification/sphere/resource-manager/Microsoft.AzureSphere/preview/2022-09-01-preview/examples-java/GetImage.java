/** Samples for Images Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetImage.json
     */
    /**
     * Sample code: Images_Get.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void imagesGet(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .images()
            .getWithResponse("MyResourceGroup1", "MyCatalog1", "myImageId", com.azure.core.util.Context.NONE);
    }
}
