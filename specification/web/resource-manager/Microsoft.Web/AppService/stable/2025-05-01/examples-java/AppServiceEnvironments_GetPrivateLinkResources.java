
/**
 * Samples for AppServiceEnvironments GetPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_GetPrivateLinkResources.json
     */
    /**
     * Sample code: Gets the private link resources.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getsThePrivateLinkResources(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getPrivateLinkResourcesWithResponse("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
