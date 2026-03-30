
/**
 * Samples for WebApps GetInstanceInfo.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSiteInstanceInfo.json
     */
    /**
     * Sample code: Get site instance info.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getSiteInstanceInfo(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getInstanceInfoWithResponse("testrg123", "tests346", "134987120",
            com.azure.core.util.Context.NONE);
    }
}
