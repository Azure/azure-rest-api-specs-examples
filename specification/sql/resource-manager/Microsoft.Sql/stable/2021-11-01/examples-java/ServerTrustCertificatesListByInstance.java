
import com.azure.core.util.Context;

/** Samples for ServerTrustCertificates ListByInstance. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerTrustCertificatesListByInstance
     * .json
     */
    /**
     * Sample code: Gets a list of server trust certificates on a given server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsAListOfServerTrustCertificatesOnAGivenServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerTrustCertificates().listByInstance("testrg", "testcl",
            Context.NONE);
    }
}
