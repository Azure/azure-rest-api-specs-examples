
/**
 * Samples for GeoCatalogs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-11-preview/GeoCatalogs_Delete.json
     */
    /**
     * Sample code: GeoCatalogs_Delete.
     * 
     * @param manager Entry point to PlanetaryComputerManager.
     */
    public static void geoCatalogsDelete(com.azure.resourcemanager.planetarycomputer.PlanetaryComputerManager manager) {
        manager.geoCatalogs().delete("MyResourceGroup", "MyCatalog", com.azure.core.util.Context.NONE);
    }
}
