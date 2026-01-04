
/**
 * Samples for AppServicePlans ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ListAppServicePlansByResourceGroup.json
     */
    /**
     * Sample code: List App Service plans by resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppServicePlansByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServicePlans().listByResourceGroup("testrg123",
            com.azure.core.util.Context.NONE);
    }
}
