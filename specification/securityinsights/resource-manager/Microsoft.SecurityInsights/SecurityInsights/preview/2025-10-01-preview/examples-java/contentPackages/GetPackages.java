
/**
 * Samples for ContentPackages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/contentPackages/GetPackages.json
     */
    /**
     * Sample code: Get all available packages.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllAvailablePackages(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.contentPackages().list("myRg", "myWorkspace", null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
