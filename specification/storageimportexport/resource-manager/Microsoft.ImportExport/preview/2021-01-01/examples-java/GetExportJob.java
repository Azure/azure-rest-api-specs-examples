/** Samples for Jobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/GetExportJob.json
     */
    /**
     * Sample code: Get export job.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void getExportJob(com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager.jobs().getByResourceGroupWithResponse("myResourceGroup", "myJob", com.azure.core.util.Context.NONE);
    }
}
