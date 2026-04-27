
/**
 * Samples for AvailableVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/AvailableVersions_ListByLocation.json
     */
    /**
     * Sample code: AvailableVersions_ListByLocation.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void availableVersionsListByLocation(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.availableVersions().listByLocation("westus2", null, com.azure.core.util.Context.NONE);
    }
}
