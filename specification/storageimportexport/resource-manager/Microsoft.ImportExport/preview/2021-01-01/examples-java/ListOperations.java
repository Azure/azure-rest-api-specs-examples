/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListOperations.json
     */
    /**
     * Sample code: List available operations.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void listAvailableOperations(
        com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
