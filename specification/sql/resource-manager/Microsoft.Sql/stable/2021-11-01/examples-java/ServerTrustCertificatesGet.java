
/**
 * Samples for ServerTrustCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerTrustCertificatesGet.json
     */
    /**
     * Sample code: Gets server trust certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsServerTrustCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerTrustCertificates().getWithResponse("testrg", "testcl",
            "customerCertificateName", com.azure.core.util.Context.NONE);
    }
}
