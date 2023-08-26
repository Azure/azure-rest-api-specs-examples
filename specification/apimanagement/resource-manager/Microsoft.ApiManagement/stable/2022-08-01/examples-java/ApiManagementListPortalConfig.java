/** Samples for PortalConfig ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListPortalConfig.json
     */
    /**
     * Sample code: ApiManagementListPortalConfig.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListPortalConfig(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.portalConfigs().listByServiceWithResponse("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
