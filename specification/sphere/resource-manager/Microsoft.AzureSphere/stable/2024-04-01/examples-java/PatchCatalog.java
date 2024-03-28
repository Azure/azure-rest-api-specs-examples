
import com.azure.resourcemanager.sphere.models.Catalog;

/**
 * Samples for Catalogs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PatchCatalog.json
     */
    /**
     * Sample code: Catalogs_Update.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsUpdate(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        Catalog resource = manager.catalogs()
            .getByResourceGroupWithResponse("MyResourceGroup1", "MyCatalog1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
