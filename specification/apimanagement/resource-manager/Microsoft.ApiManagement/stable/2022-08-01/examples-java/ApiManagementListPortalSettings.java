/** Samples for PortalSettings ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListPortalSettings.json
     */
    /**
     * Sample code: ApiManagementListPortalSettings.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListPortalSettings(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.portalSettings().listByServiceWithResponse("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
