
import com.azure.resourcemanager.sql.models.ManagedLedgerDigestUploadsName;

/**
 * Samples for ManagedLedgerDigestUploadsOperation Disable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedLedgerDigestUploadsDisable.json
     */
    /**
     * Sample code: Disables uploading ledger digests for a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        disablesUploadingLedgerDigestsForADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedLedgerDigestUploadsOperations().disable("ledgertestrg", "ledgertestserver",
            "testdb", ManagedLedgerDigestUploadsName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
