
import com.azure.resourcemanager.scvmm.models.InventoryItemProperties;

/**
 * Samples for InventoryItems Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * InventoryItems_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: InventoryItems_Create_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void inventoryItemsCreateMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.inventoryItems().define("1BdDc2Ab-bDd9-Ebd6-bfdb-C0dbbdB5DEDf").withExistingVmmServer("rgscvmm", "O")
            .withProperties(new InventoryItemProperties()).withKind("M\\d_,V.").create();
    }
}
