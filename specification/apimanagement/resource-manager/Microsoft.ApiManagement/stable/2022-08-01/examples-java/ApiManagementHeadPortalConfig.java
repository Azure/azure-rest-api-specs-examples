/** Samples for PortalConfig GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadPortalConfig.json
     */
    /**
     * Sample code: ApiManagementHeadPortalConfig.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadPortalConfig(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .portalConfigs()
            .getEntityTagWithResponse("rg1", "apimService1", "default", com.azure.core.util.Context.NONE);
    }
}
