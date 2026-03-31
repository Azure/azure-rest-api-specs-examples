
/**
 * Samples for AppServicePlans ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListAppServicePlansByResourceGroup.json
     */
    /**
     * Sample code: List App Service plans by resource group.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listAppServicePlansByResourceGroup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServicePlans().listByResourceGroup("testrg123", com.azure.core.util.Context.NONE);
    }
}
