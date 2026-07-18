
import com.azure.resourcemanager.sql.models.LedgerDigestUploadsName;

/**
 * Samples for LedgerDigestUploadsOperation Disable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/LedgerDigestUploadsDisable.json
     */
    /**
     * Sample code: Disables uploading ledger digests for a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        disablesUploadingLedgerDigestsForADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLedgerDigestUploadsOperations().disable("ledgertestrg", "ledgertestserver", "testdb",
            LedgerDigestUploadsName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
