
import com.azure.resourcemanager.sql.models.ManagedLedgerDigestUploadsName;

/**
 * Samples for ManagedLedgerDigestUploadsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedLedgerDigestUploadsGet.json
     */
    /**
     * Sample code: Gets the current ledger digest upload configuration for a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheCurrentLedgerDigestUploadConfigurationForADatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedLedgerDigestUploadsOperations().getWithResponse("ledgertestrg",
            "ledgertestserver", "testdb", ManagedLedgerDigestUploadsName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
