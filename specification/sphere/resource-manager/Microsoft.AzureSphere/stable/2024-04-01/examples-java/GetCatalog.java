
/**
 * Samples for Catalogs GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetCatalog.json
     */
    /**
     * Sample code: Catalogs_Get.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsGet(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.catalogs().getByResourceGroupWithResponse("MyResourceGroup1", "MyCatalog1",
            com.azure.core.util.Context.NONE);
    }
}
