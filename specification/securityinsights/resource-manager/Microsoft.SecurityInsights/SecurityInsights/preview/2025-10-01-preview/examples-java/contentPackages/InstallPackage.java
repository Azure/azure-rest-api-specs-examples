
import com.azure.resourcemanager.securityinsights.models.PackageKind;

/**
 * Samples for ContentPackageOperation Install.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/contentPackages/InstallPackage.json
     */
    /**
     * Sample code: Install a package to the workspace.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        installAPackageToTheWorkspace(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.contentPackageOperations().define("str.azure-sentinel-solution-str")
            .withExistingWorkspace("myRg", "myWorkspace").withContentId("str.azure-sentinel-solution-str")
            .withContentProductId("str.azure-sentinel-solution-str-sl-igl6jawr4gwmu")
            .withContentKind(PackageKind.SOLUTION).withVersion("2.0.0").withDisplayName("str").create();
    }
}
