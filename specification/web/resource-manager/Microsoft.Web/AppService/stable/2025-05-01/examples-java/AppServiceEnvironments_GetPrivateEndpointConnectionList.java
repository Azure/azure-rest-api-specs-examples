
/**
 * Samples for AppServiceEnvironments GetPrivateEndpointConnectionList.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_GetPrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets the list of private endpoints associated with a hosting environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getsTheListOfPrivateEndpointsAssociatedWithAHostingEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getPrivateEndpointConnectionList("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
