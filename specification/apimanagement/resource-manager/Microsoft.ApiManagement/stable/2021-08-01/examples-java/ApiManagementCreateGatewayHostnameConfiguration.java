/** Samples for GatewayHostnameConfiguration CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGatewayHostnameConfiguration.json
     */
    /**
     * Sample code: ApiManagementCreateGatewayHostnameConfiguration.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGatewayHostnameConfiguration(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayHostnameConfigurations()
            .define("default")
            .withExistingGateway("rg1", "apimService1", "gw1")
            .withHostname("*")
            .withCertificateId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1")
            .withNegotiateClientCertificate(false)
            .withTls10Enabled(false)
            .withTls11Enabled(false)
            .withHttp2Enabled(true)
            .create();
    }
}
