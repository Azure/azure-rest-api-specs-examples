
/**
 * Samples for EndpointCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/EndpointCertificatesGet.json
     */
    /**
     * Sample code: Gets an endpoint certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAnEndpointCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getEndpointCertificates().getWithResponse("testrg", "testcl",
            "DATABASE_MIRRORING", com.azure.core.util.Context.NONE);
    }
}
