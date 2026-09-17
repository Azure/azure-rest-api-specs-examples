
/**
 * Samples for ProductPackage Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/contentPackages/GetProductPackageById.json
     */
    /**
     * Sample code: Get a package.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAPackage(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.productPackages().getWithResponse("myRg", "myWorkspace", "str.azure-sentinel-solution-str",
            com.azure.core.util.Context.NONE);
    }
}
