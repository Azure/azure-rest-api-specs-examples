
/**
 * Samples for DeletedWebApps ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListDeletedWebAppsByLocation.json
     */
    /**
     * Sample code: List Deleted Web App by Location.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listDeletedWebAppByLocation(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDeletedWebApps().listByLocation("West US 2", com.azure.core.util.Context.NONE);
    }
}
