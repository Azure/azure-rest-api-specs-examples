import com.azure.core.util.Context;

/** Samples for Ledger GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/ConfidentialLedger_Get.json
     */
    /**
     * Sample code: ConfidentialLedgerGet.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void confidentialLedgerGet(
        com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.ledgers().getByResourceGroupWithResponse("DummyResourceGroupName", "DummyLedgerName", Context.NONE);
    }
}
