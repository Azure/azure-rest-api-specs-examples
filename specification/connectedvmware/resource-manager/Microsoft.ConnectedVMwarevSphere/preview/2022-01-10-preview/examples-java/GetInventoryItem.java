import com.azure.core.util.Context;

/** Samples for InventoryItems Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetInventoryItem.json
     */
    /**
     * Sample code: GetInventoryItem.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getInventoryItem(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.inventoryItems().getWithResponse("testrg", "ContosoVCenter", "testItem", Context.NONE);
    }
}
