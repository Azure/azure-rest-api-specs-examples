
/**
 * Samples for FileServices List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileServicesList.json
     */
    /**
     * Sample code: ListFileServices.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listFileServices(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileServices().listWithResponse("res9290", "sto1590",
            com.azure.core.util.Context.NONE);
    }
}
