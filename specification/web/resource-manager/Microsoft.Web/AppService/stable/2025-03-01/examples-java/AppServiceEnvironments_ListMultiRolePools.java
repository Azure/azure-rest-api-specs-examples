
/**
 * Samples for AppServiceEnvironments ListMultiRolePools.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_ListMultiRolePools.json
     */
    /**
     * Sample code: Get all multi-role pools.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllMultiRolePools(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listMultiRolePools("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
