
/**
 * Samples for ContentPackages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/contentPackages/GetPackageById.json
     */
    /**
     * Sample code: Get installed packages by id.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getInstalledPackagesById(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.contentPackages().getWithResponse("myRg", "myWorkspace", "str.azure-sentinel-solution-str",
            com.azure.core.util.Context.NONE);
    }
}
