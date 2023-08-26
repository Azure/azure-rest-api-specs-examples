/** Samples for PortalRevision GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadPortalRevision.json
     */
    /**
     * Sample code: ApiManagementHeadPortalRevision.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadPortalRevision(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .portalRevisions()
            .getEntityTagWithResponse("rg1", "apimService1", "20201112101010", com.azure.core.util.Context.NONE);
    }
}
