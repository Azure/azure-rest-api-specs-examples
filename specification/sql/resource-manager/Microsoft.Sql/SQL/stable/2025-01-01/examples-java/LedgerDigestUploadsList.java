
/**
 * Samples for LedgerDigestUploadsOperation ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/LedgerDigestUploadsList.json
     */
    /**
     * Sample code: Gets list of ledger digest upload settings on a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsListOfLedgerDigestUploadSettingsOnADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLedgerDigestUploadsOperations().listByDatabase("ledgertestrg", "ledgertestserver",
            "testdb", com.azure.core.util.Context.NONE);
    }
}
