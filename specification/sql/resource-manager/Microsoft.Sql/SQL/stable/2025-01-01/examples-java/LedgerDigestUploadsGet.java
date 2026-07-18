
import com.azure.resourcemanager.sql.models.LedgerDigestUploadsName;

/**
 * Samples for LedgerDigestUploadsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/LedgerDigestUploadsGet.json
     */
    /**
     * Sample code: Gets the current ledger digest upload configuration for a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheCurrentLedgerDigestUploadConfigurationForADatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLedgerDigestUploadsOperations().getWithResponse("ledgertestrg", "ledgertestserver",
            "testdb", LedgerDigestUploadsName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
