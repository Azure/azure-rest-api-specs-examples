/** Samples for GatewayCertificateAuthority CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGatewayCertificateAuthority.json
     */
    /**
     * Sample code: ApiManagementCreateGatewayCertificateAuthority.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGatewayCertificateAuthority(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayCertificateAuthorities()
            .define("cert1")
            .withExistingGateway("rg1", "apimService1", "gw1")
            .withIsTrusted(false)
            .create();
    }
}
