
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ExtendedServerBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/** Samples for ExtendedServerBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ExtendedServerBlobAuditingCreateMin.
     * json
     */
    /**
     * Sample code: Update a server's extended blob auditing policy with minimal parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAServerSExtendedBlobAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getExtendedServerBlobAuditingPolicies().createOrUpdate(
            "blobauditingtest-4799", "blobauditingtest-6440",
            new ExtendedServerBlobAuditingPolicyInner().withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder"),
            Context.NONE);
    }
}
