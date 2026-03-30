
import com.azure.resourcemanager.appservice.models.StaticSiteZipDeploymentArmResource;

/**
 * Samples for StaticSites CreateZipDeploymentForStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/StaticSiteZipDeploy.json
     */
    /**
     * Sample code: Deploy a site from a zipped package.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deployASiteFromAZippedPackage(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().createZipDeploymentForStaticSite("rg", "testStaticSite0",
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
