
/**
 * Samples for DeletedWebApps GetDeletedWebAppByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetDeletedWebAppByLocation.json
     */
    /**
     * Sample code: Get Deleted Web App by Location.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDeletedWebAppByLocation(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDeletedWebApps().getDeletedWebAppByLocationWithResponse("West US 2", "9",
            com.azure.core.util.Context.NONE);
    }
}
