
/**
 * Samples for WebApps GetSlotSiteDeploymentStatusSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetSiteDeploymentStatusSlot.json
     */
    /**
     * Sample code: Get Deployment Status Slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDeploymentStatusSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getSlotSiteDeploymentStatusSlot("rg", "testSite",
            "stage", "eacfd68b-3bbd-4ad9-99c5-98614d89c8e5", com.azure.core.util.Context.NONE);
    }
}
