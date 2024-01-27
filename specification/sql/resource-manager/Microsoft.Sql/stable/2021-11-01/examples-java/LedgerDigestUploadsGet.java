
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.LedgerDigestUploadsName;

/** Samples for LedgerDigestUploadsOperation Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/LedgerDigestUploadsGet.json
     */
    /**
     * Sample code: Gets the current ledger digest upload configuration for a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheCurrentLedgerDigestUploadConfigurationForADatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getLedgerDigestUploadsOperations().getWithResponse("ledgertestrg",
            "ledgertestserver", "testdb", LedgerDigestUploadsName.CURRENT, Context.NONE);
    }
}
