
/**
 * Samples for WebApps GetScmAllowedSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetPublishingCredentialsPolicySlot.json
     */
    /**
     * Sample code: Get SCM Allowed.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getSCMAllowed(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getScmAllowedSlotWithResponse("rg", "testSite", "stage",
            com.azure.core.util.Context.NONE);
    }
}
