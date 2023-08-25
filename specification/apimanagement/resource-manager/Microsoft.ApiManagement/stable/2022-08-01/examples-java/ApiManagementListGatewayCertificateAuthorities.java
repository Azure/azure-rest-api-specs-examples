/** Samples for GatewayCertificateAuthority ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGatewayCertificateAuthorities.json
     */
    /**
     * Sample code: ApiManagementListGatewaycertificateAuthorities.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListGatewaycertificateAuthorities(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayCertificateAuthorities()
            .listByService("rg1", "apimService1", "gw1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
