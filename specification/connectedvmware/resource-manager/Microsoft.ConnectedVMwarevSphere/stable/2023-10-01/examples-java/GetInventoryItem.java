/** Samples for InventoryItems Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/GetInventoryItem.json
     */
    /**
     * Sample code: GetInventoryItem.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getInventoryItem(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .inventoryItems()
            .getWithResponse("testrg", "ContosoVCenter", "testItem", com.azure.core.util.Context.NONE);
    }
}
