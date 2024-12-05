
import com.azure.resourcemanager.confidentialledger.models.ConfidentialLedgerBackup;

/**
 * Samples for Ledger Backup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/
     * examples/ConfidentialLedger_Backup.json
     */
    /**
     * Sample code: ConfidentialLedgerBackup.
     * 
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void
        confidentialLedgerBackup(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.ledgers().backup("DummyResourceGroupName", "DummyLedgerName",
            new ConfidentialLedgerBackup().withRestoreRegion("EastUS").withUri("DummySASUri"),
            com.azure.core.util.Context.NONE);
    }
}
