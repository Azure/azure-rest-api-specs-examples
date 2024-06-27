
/**
 * Samples for InventoryItems Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/InventoryItems_Get_MinimumSet_Gen
     * .json
     */
    /**
     * Sample code: InventoryItems_Get_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void inventoryItemsGetMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.inventoryItems().getWithResponse("rgscvmm", "_", "cacb8Ceb-efAC-bebb-ae7C-dec8C5Bb7100",
            com.azure.core.util.Context.NONE);
    }
}
