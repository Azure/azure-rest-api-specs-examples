
/**
 * Samples for WebApps ListSlotSiteDeploymentStatusesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListSiteDeploymentStatusSlot.json
     */
    /**
     * Sample code: List Deployment Status Slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listDeploymentStatusSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listSlotSiteDeploymentStatusesSlot("rg", "testSite", "stage",
            com.azure.core.util.Context.NONE);
    }
}
