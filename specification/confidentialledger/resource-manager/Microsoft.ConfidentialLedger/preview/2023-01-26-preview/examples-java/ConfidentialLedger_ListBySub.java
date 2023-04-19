/** Samples for Ledger List. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-01-26-preview/examples/ConfidentialLedger_ListBySub.json
     */
    /**
     * Sample code: ConfidentialLedgerListBySub.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void confidentialLedgerListBySub(
        com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.ledgers().list(null, com.azure.core.util.Context.NONE);
    }
}
