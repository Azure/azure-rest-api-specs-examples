
/**
 * Samples for InventoryItems ListByVmmServer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * InventoryItems_ListByVmmServer_MinimumSet_Gen.json
     */
    /**
     * Sample code: InventoryItems_ListByVmmServer_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void inventoryItemsListByVmmServerMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.inventoryItems().listByVmmServer("rgscvmm", "H", com.azure.core.util.Context.NONE);
    }
}
