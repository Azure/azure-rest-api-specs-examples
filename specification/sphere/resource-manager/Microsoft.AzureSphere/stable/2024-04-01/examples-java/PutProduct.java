
/**
 * Samples for Products CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PutProduct.json
     */
    /**
     * Sample code: Products_CreateOrUpdate.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void productsCreateOrUpdate(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.products().define("MyProduct1").withExistingCatalog("MyResourceGroup1", "MyCatalog1").create();
    }
}
