/** Samples for InventoryItems Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteInventoryItem.json
     */
    /**
     * Sample code: DeleteInventoryItem.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteInventoryItem(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .inventoryItems()
            .deleteWithResponse("testrg", "ContosoVCenter", "testItem", com.azure.core.util.Context.NONE);
    }
}
