
/**
 * Samples for WebApps ListProductionSiteDeploymentStatuses.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListSiteDeploymentStatus.json
     */
    /**
     * Sample code: List Deployment Status.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listDeploymentStatus(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listProductionSiteDeploymentStatuses("rg", "testSite",
            com.azure.core.util.Context.NONE);
    }
}
