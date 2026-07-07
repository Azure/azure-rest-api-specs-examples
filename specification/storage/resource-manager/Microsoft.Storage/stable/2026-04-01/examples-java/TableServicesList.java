
/**
 * Samples for TableServices List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/TableServicesList.json
     */
    /**
     * Sample code: TableServicesList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void tableServicesList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getTableServices().listWithResponse("res9290", "sto1590",
            com.azure.core.util.Context.NONE);
    }
}
