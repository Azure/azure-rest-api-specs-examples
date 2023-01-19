/** Samples for Jobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/GetJob.json
     */
    /**
     * Sample code: Get import job.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void getImportJob(com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager.jobs().getByResourceGroupWithResponse("myResourceGroup", "myJob", com.azure.core.util.Context.NONE);
    }
}
