
/**
 * Samples for InventoryItems Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/InventoryItems_Get_MaximumSet_Gen
     * .json
     */
    /**
     * Sample code: InventoryItems_Get_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void inventoryItemsGetMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.inventoryItems().getWithResponse("rgscvmm", "1", "2bFBede6-EEf8-becB-dBbd-B96DbBFdB3f3",
            com.azure.core.util.Context.NONE);
    }
}
