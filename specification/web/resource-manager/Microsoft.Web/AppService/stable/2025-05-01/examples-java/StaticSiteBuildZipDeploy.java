
import com.azure.resourcemanager.appservice.models.StaticSiteZipDeploymentArmResource;

/**
 * Samples for StaticSites CreateZipDeploymentForStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/StaticSiteBuildZipDeploy.json
     */
    /**
     * Sample code: Deploy a site from a zipped package to a particular static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deployASiteFromAZippedPackageToAParticularStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().createZipDeploymentForStaticSiteBuild("rg", "testStaticSite0", "12",
            new StaticSiteZipDeploymentArmResource()
                .withAppZipUrl(
                    "https://[examplestorageaccount].com/happy-sea-15afae3e-master-81828877/app-zipdeploy.zip")
                .withApiZipUrl(
                    "https://[examplestorageaccount].com/happy-sea-15afae3e-master-81828877/api-zipdeploy.zip")
                .withDeploymentTitle("Update index.html").withProvider("testProvider")
                .withFunctionLanguage("testFunctionLanguage"),
            com.azure.core.util.Context.NONE);
    }
}
