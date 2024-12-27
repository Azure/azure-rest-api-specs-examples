
/**
 * Samples for Ledger ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/
     * ConfidentialLedger_List.json
     */
    /**
     * Sample code: ConfidentialLedgerList.
     * 
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void
        confidentialLedgerList(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.ledgers().listByResourceGroup("DummyResourceGroupName", null, com.azure.core.util.Context.NONE);
    }
}
