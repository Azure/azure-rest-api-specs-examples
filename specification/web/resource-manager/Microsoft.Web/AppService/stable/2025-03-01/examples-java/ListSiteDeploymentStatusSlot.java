
/**
 * Samples for WebApps ListSlotSiteDeploymentStatusesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ListSiteDeploymentStatusSlot.json
     */
    /**
     * Sample code: List Deployment Status Slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDeploymentStatusSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listSlotSiteDeploymentStatusesSlot("rg", "testSite",
            "stage", com.azure.core.util.Context.NONE);
    }
}
