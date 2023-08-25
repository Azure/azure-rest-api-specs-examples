/** Samples for PortalRevision CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreatePortalRevision.json
     */
    /**
     * Sample code: ApiManagementCreatePortalRevision.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreatePortalRevision(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .portalRevisions()
            .define("20201112101010")
            .withExistingService("rg1", "apimService1")
            .withDescription("portal revision 1")
            .withIsCurrent(true)
            .create();
    }
}
