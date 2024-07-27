
/**
 * Samples for StorageTaskAssignmentsInstancesReport List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/storageTaskAssignmentsList/
     * ListStorageTaskAssignmentsInstancesReportSummary.json
     */
    /**
     * Sample code: ListStorageTaskAssignmentsInstancesReportSummary.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listStorageTaskAssignmentsInstancesReportSummary(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageTaskAssignmentsInstancesReports().list("res4228",
            "sto4445", null, null, com.azure.core.util.Context.NONE);
    }
}
