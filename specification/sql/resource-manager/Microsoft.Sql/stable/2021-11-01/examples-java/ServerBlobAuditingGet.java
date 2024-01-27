
import com.azure.core.util.Context;

/** Samples for ServerBlobAuditingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerBlobAuditingGet.json
     */
    /**
     * Sample code: Get a server's blob auditing policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAServerSBlobAuditingPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerBlobAuditingPolicies()
            .getWithResponse("blobauditingtest-4799", "blobauditingtest-6440", Context.NONE);
    }
}
