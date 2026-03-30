
/**
 * Samples for AppServiceEnvironments Upgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_Upgrade.json
     */
    /**
     * Sample code: Initiate an upgrade on an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        initiateAnUpgradeOnAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().upgrade("rg", "SampleHostingEnvironment",
            com.azure.core.util.Context.NONE);
    }
}
