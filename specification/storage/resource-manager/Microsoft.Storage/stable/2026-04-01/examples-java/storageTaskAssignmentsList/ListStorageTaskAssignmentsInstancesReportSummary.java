
/**
 * Samples for StorageTaskAssignmentsInstancesReport List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/storageTaskAssignmentsList/ListStorageTaskAssignmentsInstancesReportSummary.json
     */
    /**
     * Sample code: ListStorageTaskAssignmentsInstancesReportSummary.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        listStorageTaskAssignmentsInstancesReportSummary(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageTaskAssignmentsInstancesReports().list("res4228", "sto4445", null, null,
            com.azure.core.util.Context.NONE);
    }
}
