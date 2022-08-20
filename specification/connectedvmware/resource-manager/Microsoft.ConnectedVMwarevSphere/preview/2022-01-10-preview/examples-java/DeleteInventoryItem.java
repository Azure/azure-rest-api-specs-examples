import com.azure.core.util.Context;

/** Samples for InventoryItems Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteInventoryItem.json
     */
    /**
     * Sample code: DeleteInventoryItem.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteInventoryItem(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.inventoryItems().deleteWithResponse("testrg", "ContosoVCenter", "testItem", Context.NONE);
    }
}
