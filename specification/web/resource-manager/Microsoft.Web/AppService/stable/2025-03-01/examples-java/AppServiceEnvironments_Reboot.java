
/**
 * Samples for AppServiceEnvironments Reboot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_Reboot.json
     */
    /**
     * Sample code: Reboot all machines in an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        rebootAllMachinesInAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().rebootWithResponse("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
