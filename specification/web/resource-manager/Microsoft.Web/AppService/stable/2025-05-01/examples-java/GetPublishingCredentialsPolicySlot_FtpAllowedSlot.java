
/**
 * Samples for WebApps GetFtpAllowedSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetPublishingCredentialsPolicySlot_FtpAllowedSlot.json
     */
    /**
     * Sample code: Get FTP Allowed.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getFTPAllowed(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getFtpAllowedSlotWithResponse("rg", "testSite", "stage",
            com.azure.core.util.Context.NONE);
    }
}
