
/**
 * Samples for InventoryItems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * InventoryItems_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: InventoryItems_Delete_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void inventoryItemsDeleteMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.inventoryItems().deleteWithResponse("rgscvmm", "b", "EcECadfd-Eaaa-e5Ce-ebdA-badeEd3c6af1",
            com.azure.core.util.Context.NONE);
    }
}
