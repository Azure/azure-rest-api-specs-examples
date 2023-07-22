/** Samples for Products Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/DeleteProduct.json
     */
    /**
     * Sample code: Products_Delete.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void productsDelete(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.products().delete("MyResourceGroup1", "MyCatalog1", "MyProductName1", com.azure.core.util.Context.NONE);
    }
}
