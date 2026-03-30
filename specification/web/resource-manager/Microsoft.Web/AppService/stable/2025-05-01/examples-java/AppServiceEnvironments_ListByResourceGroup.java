
/**
 * Samples for AppServiceEnvironments ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListByResourceGroup.json
     */
    /**
     * Sample code: Get all App Service Environments in a resource group.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAllAppServiceEnvironmentsInAResourceGroup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listByResourceGroup("test-rg",
            com.azure.core.util.Context.NONE);
    }
}
