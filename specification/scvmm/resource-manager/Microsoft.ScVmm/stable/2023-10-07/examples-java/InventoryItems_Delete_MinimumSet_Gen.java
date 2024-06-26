
/**
 * Samples for InventoryItems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * InventoryItems_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: InventoryItems_Delete_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void inventoryItemsDeleteMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.inventoryItems().deleteWithResponse("rgscvmm", "_", "cDBcbae6-BC3d-52fe-CedC-7eFeaBFabb82",
            com.azure.core.util.Context.NONE);
    }
}
