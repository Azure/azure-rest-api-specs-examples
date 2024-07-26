
/**
 * Samples for AppServiceEnvironments ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/
     * AppServiceEnvironments_ListByResourceGroup.json
     */
    /**
     * Sample code: Get all App Service Environments in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAllAppServiceEnvironmentsInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listByResourceGroup("test-rg",
            com.azure.core.util.Context.NONE);
    }
}
