
/**
 * Samples for Products ListByCatalog.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetProducts.json
     */
    /**
     * Sample code: Products_ListByCatalog.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void productsListByCatalog(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.products().listByCatalog("MyResourceGroup1", "MyCatalog1", com.azure.core.util.Context.NONE);
    }
}
