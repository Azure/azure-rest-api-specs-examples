import com.azure.resourcemanager.sphere.models.Product;

/** Samples for Products Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PatchProduct.json
     */
    /**
     * Sample code: Products_Update.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void productsUpdate(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        Product resource =
            manager
                .products()
                .getWithResponse("MyResourceGroup1", "MyCatalog1", "MyProduct1", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
