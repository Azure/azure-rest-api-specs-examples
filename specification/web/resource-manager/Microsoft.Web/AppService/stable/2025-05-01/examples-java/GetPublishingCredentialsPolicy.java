
/**
 * Samples for WebApps GetScmAllowed.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetPublishingCredentialsPolicy.json
     */
    /**
     * Sample code: Get SCM Allowed.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getSCMAllowed(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getScmAllowedWithResponse("rg", "testSite",
            com.azure.core.util.Context.NONE);
    }
}
