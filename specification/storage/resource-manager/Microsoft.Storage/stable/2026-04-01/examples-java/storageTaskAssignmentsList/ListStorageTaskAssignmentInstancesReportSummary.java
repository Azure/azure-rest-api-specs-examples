
/**
 * Samples for StorageTaskAssignmentInstancesReport List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/storageTaskAssignmentsList/ListStorageTaskAssignmentInstancesReportSummary.json
     */
    /**
     * Sample code: ListStorageTaskAssignmentInstancesReportSummary.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        listStorageTaskAssignmentInstancesReportSummary(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageTaskAssignmentInstancesReports().list("res4228", "sto4445", "myassignment1",
            null, null, com.azure.core.util.Context.NONE);
    }
}
