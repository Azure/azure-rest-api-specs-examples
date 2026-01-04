
/**
 * Samples for WebApps ListProductionSiteDeploymentStatuses.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListSiteDeploymentStatus.
     * json
     */
    /**
     * Sample code: List Deployment Status.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDeploymentStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listProductionSiteDeploymentStatuses("rg", "testSite",
            com.azure.core.util.Context.NONE);
    }
}
