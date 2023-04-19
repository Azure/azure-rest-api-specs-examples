/** Samples for Ledger Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-01-26-preview/examples/ConfidentialLedger_Delete.json
     */
    /**
     * Sample code: ConfidentialLedgerDelete.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void confidentialLedgerDelete(
        com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.ledgers().delete("DummyResourceGroupName", "DummyLedgerName", com.azure.core.util.Context.NONE);
    }
}
