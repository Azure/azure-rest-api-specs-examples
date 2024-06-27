
/**
 * Samples for InventoryItems Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * InventoryItems_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: InventoryItems_Create_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void inventoryItemsCreateMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.inventoryItems().define("bbFb0cBb-50ce-4bfc-3eeD-bC26AbCC257a").withExistingVmmServer("rgscvmm", ".")
            .create();
    }
}
