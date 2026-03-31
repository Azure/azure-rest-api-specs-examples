
/**
 * Samples for AppServiceEnvironments DeletePrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_DeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: Deletes a private endpoint connection.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        deletesAPrivateEndpointConnection(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().deletePrivateEndpointConnection("test-rg", "test-ase",
            "fa38656c-034e-43d8-adce-fe06ce039c98", com.azure.core.util.Context.NONE);
    }
}
