
/**
 * Samples for AppServiceEnvironments ListOperations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/AppServiceEnvironments_ListOperations
     * .json
     */
    /**
     * Sample code: List all currently running operations on the App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllCurrentlyRunningOperationsOnTheAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listOperationsWithResponse("test-rg",
            "test-ase", com.azure.core.util.Context.NONE);
    }
}
