
/**
 * Samples for WebApps GetProductionSiteDeploymentStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetSiteDeploymentStatus.json
     */
    /**
     * Sample code: Get Deployment Status.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDeploymentStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getProductionSiteDeploymentStatus("rg", "testSite",
            "eacfd68b-3bbd-4ad9-99c5-98614d89c8e5", com.azure.core.util.Context.NONE);
    }
}
