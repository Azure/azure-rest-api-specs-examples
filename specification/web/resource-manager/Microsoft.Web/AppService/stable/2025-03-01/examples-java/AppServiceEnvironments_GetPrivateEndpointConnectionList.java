
/**
 * Samples for AppServiceEnvironments GetPrivateEndpointConnectionList.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_GetPrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets the list of private endpoints associated with a hosting environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheListOfPrivateEndpointsAssociatedWithAHostingEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments()
            .getPrivateEndpointConnectionList("test-rg", "test-ase", com.azure.core.util.Context.NONE);
    }
}
