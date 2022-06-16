/** Samples for InventoryItems Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateInventoryItem.json
     */
    /**
     * Sample code: CreateInventoryItem.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void createInventoryItem(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .inventoryItems()
            .define("12345678-1234-1234-1234-123456789abc")
            .withExistingVmmServer("testrg", "ContosoVMMServer")
            .create();
    }
}
