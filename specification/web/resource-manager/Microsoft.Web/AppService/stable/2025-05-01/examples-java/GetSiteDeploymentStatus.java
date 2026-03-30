
/**
 * Samples for WebApps GetProductionSiteDeploymentStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSiteDeploymentStatus.json
     */
    /**
     * Sample code: Get Deployment Status.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDeploymentStatus(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getProductionSiteDeploymentStatus("rg", "testSite",
            "eacfd68b-3bbd-4ad9-99c5-98614d89c8e5", com.azure.core.util.Context.NONE);
    }
}
