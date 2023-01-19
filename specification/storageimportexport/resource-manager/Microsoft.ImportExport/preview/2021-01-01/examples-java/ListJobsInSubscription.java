/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListJobsInSubscription.json
     */
    /**
     * Sample code: List jobs in a subscription.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void listJobsInASubscription(
        com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager.jobs().list(null, null, com.azure.core.util.Context.NONE);
    }
}
