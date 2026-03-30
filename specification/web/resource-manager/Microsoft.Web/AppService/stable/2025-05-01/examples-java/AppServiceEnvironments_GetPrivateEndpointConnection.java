
/**
 * Samples for AppServiceEnvironments GetPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_GetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Gets a private endpoint connection.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getsAPrivateEndpointConnection(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getPrivateEndpointConnectionWithResponse("test-rg",
            "test-ase", "fa38656c-034e-43d8-adce-fe06ce039c98", com.azure.core.util.Context.NONE);
    }
}
