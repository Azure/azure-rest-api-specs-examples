
/**
 * Samples for GeoCatalogs GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-11-preview/GeoCatalogs_Get.json
     */
    /**
     * Sample code: GeoCatalogs_Get.
     * 
     * @param manager Entry point to PlanetaryComputerManager.
     */
    public static void geoCatalogsGet(com.azure.resourcemanager.planetarycomputer.PlanetaryComputerManager manager) {
        manager.geoCatalogs().getByResourceGroupWithResponse("MyResourceGroup", "MyCatalog",
            com.azure.core.util.Context.NONE);
    }
}
