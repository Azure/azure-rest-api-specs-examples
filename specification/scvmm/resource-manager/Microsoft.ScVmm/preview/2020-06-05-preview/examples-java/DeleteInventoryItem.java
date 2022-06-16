import com.azure.core.util.Context;

/** Samples for InventoryItems Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteInventoryItem.json
     */
    /**
     * Sample code: DeleteInventoryItem.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void deleteInventoryItem(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .inventoryItems()
            .deleteWithResponse("testrg", "ContosoVMMServer", "12345678-1234-1234-1234-123456789abc", Context.NONE);
    }
}
