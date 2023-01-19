/** Samples for Locations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListLocations.json
     */
    /**
     * Sample code: List locations.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void listLocations(com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager.locations().list(com.azure.core.util.Context.NONE);
    }
}
