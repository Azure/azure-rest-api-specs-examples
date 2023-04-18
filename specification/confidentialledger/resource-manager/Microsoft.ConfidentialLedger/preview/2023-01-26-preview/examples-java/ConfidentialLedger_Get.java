/** Samples for Ledger GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-01-26-preview/examples/ConfidentialLedger_Get.json
     */
    /**
     * Sample code: ConfidentialLedgerGet.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void confidentialLedgerGet(
        com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager
            .ledgers()
            .getByResourceGroupWithResponse(
                "DummyResourceGroupName", "DummyLedgerName", com.azure.core.util.Context.NONE);
    }
}
