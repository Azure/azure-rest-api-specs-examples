
/**
 * Samples for StorageTasksReport List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/
     * storageTasksList/ListStorageTasksRunReportSummary.json
     */
    /**
     * Sample code: ListStorageTasksByResourceGroup.
     * 
     * @param manager Entry point to StorageActionsManager.
     */
    public static void
        listStorageTasksByResourceGroup(com.azure.resourcemanager.storageactions.StorageActionsManager manager) {
        manager.storageTasksReports().list("rgroup1", "mytask1", null, null, com.azure.core.util.Context.NONE);
    }
}
