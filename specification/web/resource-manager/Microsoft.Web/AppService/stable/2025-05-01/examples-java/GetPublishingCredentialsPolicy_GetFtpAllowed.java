
/**
 * Samples for WebApps GetFtpAllowed.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetPublishingCredentialsPolicy_GetFtpAllowed.json
     */
    /**
     * Sample code: Get FTP Allowed.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getFTPAllowed(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getFtpAllowedWithResponse("rg", "testSite",
            com.azure.core.util.Context.NONE);
    }
}
