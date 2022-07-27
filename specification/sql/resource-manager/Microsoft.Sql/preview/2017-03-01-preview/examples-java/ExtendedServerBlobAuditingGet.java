import com.azure.core.util.Context;

/** Samples for ExtendedServerBlobAuditingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ExtendedServerBlobAuditingGet.json
     */
    /**
     * Sample code: Get a server's blob extended auditing policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAServerSBlobExtendedAuditingPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getExtendedServerBlobAuditingPolicies()
            .getWithResponse("blobauditingtest-4799", "blobauditingtest-6440", Context.NONE);
    }
}
