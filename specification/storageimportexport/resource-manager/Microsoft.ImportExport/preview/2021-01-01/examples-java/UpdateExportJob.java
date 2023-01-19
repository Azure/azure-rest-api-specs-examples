import com.azure.resourcemanager.storageimportexport.models.JobResponse;

/** Samples for Jobs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/UpdateExportJob.json
     */
    /**
     * Sample code: Update export job.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void updateExportJob(
        com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        JobResponse resource =
            manager
                .jobs()
                .getByResourceGroupWithResponse("myResourceGroup", "myExportJob", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withState("").withLogLevel("Verbose").withBackupDriveManifest(true).apply();
    }
}
