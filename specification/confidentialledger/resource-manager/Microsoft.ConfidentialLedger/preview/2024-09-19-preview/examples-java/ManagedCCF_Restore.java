
import com.azure.resourcemanager.confidentialledger.models.ManagedCcfRestore;

/**
 * Samples for ManagedCcf Restore.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2024-09-19-preview/
     * examples/ManagedCCF_Restore.json
     */
    /**
     * Sample code: ManagedCCFRestore.
     * 
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void
        managedCCFRestore(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.managedCcfs().restore(
            "DummyResourceGroupName", "DummyMccfAppName", new ManagedCcfRestore()
                .withFileShareName("DummyFileShareName").withRestoreRegion("EastUS").withUri("DummySASUri"),
            com.azure.core.util.Context.NONE);
    }
}
