
import com.azure.core.util.Context;

/** Samples for LedgerDigestUploadsOperation ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/LedgerDigestUploadsList.json
     */
    /**
     * Sample code: Gets list of ledger digest upload settings on a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsListOfLedgerDigestUploadSettingsOnADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getLedgerDigestUploadsOperations().listByDatabase("ledgertestrg",
            "ledgertestserver", "testdb", Context.NONE);
    }
}
