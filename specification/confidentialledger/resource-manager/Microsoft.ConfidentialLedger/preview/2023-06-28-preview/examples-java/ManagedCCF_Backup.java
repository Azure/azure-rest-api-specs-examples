
import com.azure.resourcemanager.confidentialledger.models.ManagedCcfBackup;

/**
 * Samples for ManagedCcf Backup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/
     * examples/ManagedCCF_Backup.json
     */
    /**
     * Sample code: ManagedCCFBackup.
     * 
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void
        managedCCFBackup(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.managedCcfs().backup("DummyResourceGroupName", "DummyMccfAppName",
            new ManagedCcfBackup().withRestoreRegion("EastUS").withUri("DummySASUri"),
            com.azure.core.util.Context.NONE);
    }
}
