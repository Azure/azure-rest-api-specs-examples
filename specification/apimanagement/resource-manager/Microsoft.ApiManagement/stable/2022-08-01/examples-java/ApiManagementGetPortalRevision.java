/** Samples for PortalRevision Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetPortalRevision.json
     */
    /**
     * Sample code: ApiManagementGetPortalRevision.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetPortalRevision(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .portalRevisions()
            .getWithResponse("rg1", "apimService1", "20201112101010", com.azure.core.util.Context.NONE);
    }
}
