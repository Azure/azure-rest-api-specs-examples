/** Samples for AppServiceEnvironments Upgrade. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/AppServiceEnvironments_Upgrade.json
     */
    /**
     * Sample code: Initiate an upgrade on an App Service Environment.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void initiateAnUpgradeOnAnAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .upgrade("rg", "SampleHostingEnvironment", com.azure.core.util.Context.NONE);
    }
}
