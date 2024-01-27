
import com.azure.core.util.Context;

/** Samples for EndpointCertificates ListByInstance. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/EndpointCertificatesListByInstance.
     * json
     */
    /**
     * Sample code: Get a list of endpoint certificates.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAListOfEndpointCertificates(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getEndpointCertificates().listByInstance("testrg", "testcl",
            Context.NONE);
    }
}
