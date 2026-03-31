
/**
 * Samples for Global GetDeletedWebAppSnapshots.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetDeletedWebAppSnapshots.json
     */
    /**
     * Sample code: Get Deleted Web App Snapshots.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDeletedWebAppSnapshots(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getGlobals().getDeletedWebAppSnapshotsWithResponse("9",
            com.azure.core.util.Context.NONE);
    }
}
