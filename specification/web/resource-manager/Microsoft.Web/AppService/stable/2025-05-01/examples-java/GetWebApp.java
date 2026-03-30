
/**
 * Samples for WebApps GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetWebApp.json
     */
    /**
     * Sample code: Get Web App.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getWebApp(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getByResourceGroupWithResponse("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
