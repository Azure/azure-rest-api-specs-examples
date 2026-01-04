
/**
 * Samples for AppServiceEnvironments GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/AppServiceEnvironments_Get
     * .json
     */
    /**
     * Sample code: Get the properties of an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getThePropertiesOfAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().getByResourceGroupWithResponse("test-rg",
            "test-ase", com.azure.core.util.Context.NONE);
    }
}
