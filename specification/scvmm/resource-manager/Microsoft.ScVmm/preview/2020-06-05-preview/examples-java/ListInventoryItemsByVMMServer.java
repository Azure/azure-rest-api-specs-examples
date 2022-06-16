import com.azure.core.util.Context;

/** Samples for InventoryItems ListByVmmServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListInventoryItemsByVMMServer.json
     */
    /**
     * Sample code: InventoryItemsListByVMMServer.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void inventoryItemsListByVMMServer(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.inventoryItems().listByVmmServer("testrg", "ContosoVMMServer", Context.NONE);
    }
}
