
/**
 * Samples for GeoCatalogs ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-11-preview/GeoCatalogs_ListByResourceGroup.json
     */
    /**
     * Sample code: GeoCatalogs_ListByResourceGroup.
     * 
     * @param manager Entry point to PlanetaryComputerManager.
     */
    public static void
        geoCatalogsListByResourceGroup(com.azure.resourcemanager.planetarycomputer.PlanetaryComputerManager manager) {
        manager.geoCatalogs().listByResourceGroup("MyResourceGroup", com.azure.core.util.Context.NONE);
    }
}
