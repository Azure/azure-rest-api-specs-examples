
import com.azure.resourcemanager.sql.models.LedgerDigestUploadsName;

/**
 * Samples for LedgerDigestUploadsOperation Disable.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/LedgerDigestUploadsDisable.json
     */
    /**
     * Sample code: Disables uploading ledger digests for a database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        disablesUploadingLedgerDigestsForADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getLedgerDigestUploadsOperations().disable("ledgertestrg",
            "ledgertestserver", "testdb", LedgerDigestUploadsName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
