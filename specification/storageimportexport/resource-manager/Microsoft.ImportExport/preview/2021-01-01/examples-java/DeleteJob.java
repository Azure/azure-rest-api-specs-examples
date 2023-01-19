/** Samples for Jobs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/DeleteJob.json
     */
    /**
     * Sample code: Delete job.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void deleteJob(com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager.jobs().deleteByResourceGroupWithResponse("myResourceGroup", "myJob", com.azure.core.util.Context.NONE);
    }
}
