/** Samples for GatewayCertificateAuthority Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetGatewayCertificateAuthority.json
     */
    /**
     * Sample code: ApiManagementGetGatewayCertificateAuthority.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetGatewayCertificateAuthority(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayCertificateAuthorities()
            .getWithResponse("rg1", "apimService1", "gw1", "cert1", com.azure.core.util.Context.NONE);
    }
}
