import com.azure.core.util.Context;

/** Samples for Ledger List. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/ConfidentialLedger_ListBySub.json
     */
    /**
     * Sample code: ConfidentialLedgerListBySub.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void confidentialLedgerListBySub(
        com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.ledgers().list(null, Context.NONE);
    }
}
