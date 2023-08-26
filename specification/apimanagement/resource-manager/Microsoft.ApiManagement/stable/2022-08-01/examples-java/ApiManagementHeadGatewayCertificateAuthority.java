/** Samples for GatewayCertificateAuthority GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadGatewayCertificateAuthority.json
     */
    /**
     * Sample code: ApiManagementHeadGatewayCertificateAuthority.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadGatewayCertificateAuthority(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayCertificateAuthorities()
            .getEntityTagWithResponse("rg1", "apimService1", "gw1", "cert1", com.azure.core.util.Context.NONE);
    }
}
