
/**
 * Samples for AvailableVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/AvailableVersions_ListByLocationWithFilter.json
     */
    /**
     * Sample code: AvailableVersions_ListByLocationWithFilter.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void
        availableVersionsListByLocationWithFilter(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.availableVersions().listByLocation("westus2", "1.28", com.azure.core.util.Context.NONE);
    }
}
