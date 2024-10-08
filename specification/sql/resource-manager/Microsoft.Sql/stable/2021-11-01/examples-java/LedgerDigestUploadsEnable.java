
import com.azure.resourcemanager.sql.fluent.models.LedgerDigestUploadsInner;
import com.azure.resourcemanager.sql.models.LedgerDigestUploadsName;

/**
 * Samples for LedgerDigestUploadsOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/LedgerDigestUploadsEnable.json
     */
    /**
     * Sample code: Enables ledger digest upload configuration for a database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        enablesLedgerDigestUploadConfigurationForADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getLedgerDigestUploadsOperations().createOrUpdate("ledgertestrg",
            "ledgertestserver", "testdb", LedgerDigestUploadsName.CURRENT,
            new LedgerDigestUploadsInner().withDigestStorageEndpoint("https://MyAccount.blob.core.windows.net"),
            com.azure.core.util.Context.NONE);
    }
}
