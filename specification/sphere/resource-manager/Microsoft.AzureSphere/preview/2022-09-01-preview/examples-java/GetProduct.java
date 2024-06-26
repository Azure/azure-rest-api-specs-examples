/** Samples for Products Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetProduct.json
     */
    /**
     * Sample code: Products_Get.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void productsGet(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .products()
            .getWithResponse("MyResourceGroup1", "MyCatalog1", "MyProduct1", com.azure.core.util.Context.NONE);
    }
}
