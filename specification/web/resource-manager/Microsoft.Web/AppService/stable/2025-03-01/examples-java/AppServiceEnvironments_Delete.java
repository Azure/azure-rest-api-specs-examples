
/**
 * Samples for AppServiceEnvironments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_Delete.json
     */
    /**
     * Sample code: Delete an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().delete("test-rg", "test-ase", null,
            com.azure.core.util.Context.NONE);
    }
}
