/** Samples for BitLockerKeys List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListBitLockerKeys.json
     */
    /**
     * Sample code: List BitLocker Keys for drives in a job.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void listBitLockerKeysForDrivesInAJob(
        com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager.bitLockerKeys().list("myJob", "myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
