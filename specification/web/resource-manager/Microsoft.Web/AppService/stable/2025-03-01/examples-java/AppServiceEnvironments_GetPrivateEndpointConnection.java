
/**
 * Samples for AppServiceEnvironments GetPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_GetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Gets a private endpoint connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().getPrivateEndpointConnectionWithResponse(
            "test-rg", "test-ase", "fa38656c-034e-43d8-adce-fe06ce039c98", com.azure.core.util.Context.NONE);
    }
}
