
/**
 * Samples for InventoryItems ListByVmmServer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * InventoryItems_ListByVmmServer_MaximumSet_Gen.json
     */
    /**
     * Sample code: InventoryItems_ListByVmmServer_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void inventoryItemsListByVmmServerMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.inventoryItems().listByVmmServer("rgscvmm", "X", com.azure.core.util.Context.NONE);
    }
}
