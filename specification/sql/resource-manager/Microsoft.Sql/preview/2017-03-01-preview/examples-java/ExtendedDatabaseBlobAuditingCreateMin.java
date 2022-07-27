import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ExtendedDatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/** Samples for ExtendedDatabaseBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ExtendedDatabaseBlobAuditingCreateMin.json
     */
    /**
     * Sample code: Create or update an extended database's blob auditing policy with minimal parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAnExtendedDatabaseSBlobAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getExtendedDatabaseBlobAuditingPolicies()
            .createOrUpdateWithResponse(
                "blobauditingtest-4799",
                "blobauditingtest-6440",
                "testdb",
                new ExtendedDatabaseBlobAuditingPolicyInner()
                    .withState(BlobAuditingPolicyState.ENABLED)
                    .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                    .withStorageAccountAccessKey(
                        "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
                Context.NONE);
    }
}
