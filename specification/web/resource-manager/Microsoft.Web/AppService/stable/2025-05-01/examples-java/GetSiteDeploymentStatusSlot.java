
/**
 * Samples for WebApps GetSlotSiteDeploymentStatusSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSiteDeploymentStatusSlot.json
     */
    /**
     * Sample code: Get Deployment Status Slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDeploymentStatusSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getSlotSiteDeploymentStatusSlot("rg", "testSite", "stage",
            "eacfd68b-3bbd-4ad9-99c5-98614d89c8e5", com.azure.core.util.Context.NONE);
    }
}
