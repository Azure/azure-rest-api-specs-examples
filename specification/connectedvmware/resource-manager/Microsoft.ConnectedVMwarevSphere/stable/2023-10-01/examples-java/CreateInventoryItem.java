import com.azure.resourcemanager.connectedvmware.models.ResourcePoolInventoryItem;

/** Samples for InventoryItems Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/CreateInventoryItem.json
     */
    /**
     * Sample code: CreateInventoryItem.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void createInventoryItem(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .inventoryItems()
            .define("testItem")
            .withExistingVcenter("testrg", "ContosoVCenter")
            .withProperties(new ResourcePoolInventoryItem())
            .create();
    }
}
