
/**
 * Samples for GeoCatalogs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-11-preview/GeoCatalogs_ListBySubscription.json
     */
    /**
     * Sample code: GeoCatalogs_ListBySubscription.
     * 
     * @param manager Entry point to PlanetaryComputerManager.
     */
    public static void
        geoCatalogsListBySubscription(com.azure.resourcemanager.planetarycomputer.PlanetaryComputerManager manager) {
        manager.geoCatalogs().list(com.azure.core.util.Context.NONE);
    }
}
