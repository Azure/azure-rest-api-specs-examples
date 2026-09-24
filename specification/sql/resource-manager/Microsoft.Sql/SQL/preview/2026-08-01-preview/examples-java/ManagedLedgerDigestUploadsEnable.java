
import com.azure.resourcemanager.sql.fluent.models.ManagedLedgerDigestUploadsInner;
import com.azure.resourcemanager.sql.models.ManagedLedgerDigestUploadsName;

/**
 * Samples for ManagedLedgerDigestUploadsOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedLedgerDigestUploadsEnable.json
     */
    /**
     * Sample code: Enables managed ledger digest upload configuration for a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void enablesManagedLedgerDigestUploadConfigurationForADatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedLedgerDigestUploadsOperations().createOrUpdate("ledgertestrg",
            "ledgertestserver", "testdb", ManagedLedgerDigestUploadsName.CURRENT,
            new ManagedLedgerDigestUploadsInner().withDigestStorageEndpoint("https://MyAccount.blob.core.windows.net"),
            com.azure.core.util.Context.NONE);
    }
}
