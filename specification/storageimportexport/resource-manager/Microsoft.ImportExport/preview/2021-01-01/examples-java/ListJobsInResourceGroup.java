/** Samples for Jobs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListJobsInResourceGroup.json
     */
    /**
     * Sample code: List jobs in a resource group.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void listJobsInAResourceGroup(
        com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager.jobs().listByResourceGroup("myResourceGroup", null, null, com.azure.core.util.Context.NONE);
    }
}
