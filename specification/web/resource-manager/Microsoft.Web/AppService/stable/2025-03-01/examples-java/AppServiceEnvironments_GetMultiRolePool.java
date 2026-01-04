
/**
 * Samples for AppServiceEnvironments GetMultiRolePool.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_GetMultiRolePool.json
     */
    /**
     * Sample code: Get properties of a multi-role pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPropertiesOfAMultiRolePool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().getMultiRolePoolWithResponse("test-rg",
            "test-ase", com.azure.core.util.Context.NONE);
    }
}
