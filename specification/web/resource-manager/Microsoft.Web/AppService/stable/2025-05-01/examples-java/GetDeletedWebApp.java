
/**
 * Samples for Global GetDeletedWebApp.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetDeletedWebApp.json
     */
    /**
     * Sample code: Get Deleted Web App.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDeletedWebApp(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getGlobals().getDeletedWebAppWithResponse("9", com.azure.core.util.Context.NONE);
    }
}
