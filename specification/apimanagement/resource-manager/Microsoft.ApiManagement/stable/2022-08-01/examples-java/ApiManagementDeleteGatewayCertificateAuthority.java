/** Samples for GatewayCertificateAuthority Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteGatewayCertificateAuthority.json
     */
    /**
     * Sample code: ApiManagementDeleteGatewayCertificateAuthority.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteGatewayCertificateAuthority(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayCertificateAuthorities()
            .deleteWithResponse("rg1", "apimService1", "gw1", "default", "*", com.azure.core.util.Context.NONE);
    }
}
