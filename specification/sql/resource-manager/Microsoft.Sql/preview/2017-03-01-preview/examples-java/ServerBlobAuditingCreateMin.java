import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/** Samples for ServerBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ServerBlobAuditingCreateMin.json
     */
    /**
     * Sample code: Update a server's blob auditing policy with minimal parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAServerSBlobAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerBlobAuditingPolicies()
            .createOrUpdate(
                "blobauditingtest-4799",
                "blobauditingtest-6440",
                new ServerBlobAuditingPolicyInner()
                    .withState(BlobAuditingPolicyState.ENABLED)
                    .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                    .withStorageAccountAccessKey(
                        "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
                Context.NONE);
    }
}
