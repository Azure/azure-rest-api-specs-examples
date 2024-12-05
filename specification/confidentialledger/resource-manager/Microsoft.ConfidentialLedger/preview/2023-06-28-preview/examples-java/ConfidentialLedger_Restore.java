
import com.azure.resourcemanager.confidentialledger.models.ConfidentialLedgerRestore;

/**
 * Samples for Ledger Restore.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/
     * examples/ConfidentialLedger_Restore.json
     */
    /**
     * Sample code: ConfidentialLedgerRestore.
     * 
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void
        confidentialLedgerRestore(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.ledgers().restore(
            "DummyResourceGroupName", "DummyLedgerName", new ConfidentialLedgerRestore()
                .withFileShareName("DummyFileShareName").withRestoreRegion("EastUS").withUri("DummySASUri"),
            com.azure.core.util.Context.NONE);
    }
}
