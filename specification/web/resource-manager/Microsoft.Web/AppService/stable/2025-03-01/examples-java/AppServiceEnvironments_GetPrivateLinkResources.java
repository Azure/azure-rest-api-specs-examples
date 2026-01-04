
/**
 * Samples for AppServiceEnvironments GetPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_GetPrivateLinkResources.json
     */
    /**
     * Sample code: Gets the private link resources.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsThePrivateLinkResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments()
            .getPrivateLinkResourcesWithResponse("test-rg", "test-ase", com.azure.core.util.Context.NONE);
    }
}
