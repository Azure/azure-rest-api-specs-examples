
import com.azure.resourcemanager.appservice.models.StaticSiteZipDeploymentArmResource;

/**
 * Samples for StaticSites CreateZipDeploymentForStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/StaticSiteBuildZipDeploy.json
     */
    /**
     * Sample code: Deploy a site from a zipped package to a particular static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deployASiteFromAZippedPackageToAParticularStaticSiteBuild(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().createZipDeploymentForStaticSiteBuild("rg",
            "testStaticSite0", "12",
            new StaticSiteZipDeploymentArmResource()
                .withAppZipUrl("https://teststorageaccount.net/happy-sea-15afae3e-master-81828877/app-zipdeploy.zip")
                .withApiZipUrl("https://teststorageaccount.net/happy-sea-15afae3e-master-81828877/api-zipdeploy.zip")
                .withDeploymentTitle("Update index.html").withProvider("testProvider")
                .withFunctionLanguage("testFunctionLanguage"),
            com.azure.core.util.Context.NONE);
    }
}
