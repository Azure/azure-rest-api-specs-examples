
/**
 * Samples for AppServiceEnvironments Suspend.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/AppServiceEnvironments_Suspend.json
     */
    /**
     * Sample code: Suspend an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void suspendAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().suspend("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
