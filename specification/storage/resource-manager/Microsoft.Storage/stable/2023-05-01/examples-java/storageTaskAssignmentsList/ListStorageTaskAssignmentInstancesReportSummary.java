
/**
 * Samples for StorageTaskAssignmentInstancesReport List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/storageTaskAssignmentsList/
     * ListStorageTaskAssignmentInstancesReportSummary.json
     */
    /**
     * Sample code: ListStorageTaskAssignmentInstancesReportSummary.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listStorageTaskAssignmentInstancesReportSummary(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageTaskAssignmentInstancesReports().list("res4228",
            "sto4445", "myassignment1", null, null, com.azure.core.util.Context.NONE);
    }
}
