
import com.azure.resourcemanager.sql.fluent.models.LedgerDigestUploadsInner;
import com.azure.resourcemanager.sql.models.LedgerDigestUploadsName;

/**
 * Samples for LedgerDigestUploadsOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/LedgerDigestUploadsEnable.json
     */
    /**
     * Sample code: Enables ledger digest upload configuration for a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        enablesLedgerDigestUploadConfigurationForADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLedgerDigestUploadsOperations().createOrUpdate("ledgertestrg", "ledgertestserver",
            "testdb", LedgerDigestUploadsName.CURRENT,
            new LedgerDigestUploadsInner().withDigestStorageEndpoint("https://MyAccount.blob.core.windows.net"),
            com.azure.core.util.Context.NONE);
    }
}
