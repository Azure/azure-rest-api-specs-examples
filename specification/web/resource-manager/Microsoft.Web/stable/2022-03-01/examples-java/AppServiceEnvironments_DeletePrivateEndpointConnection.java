import com.azure.core.util.Context;

/** Samples for AppServiceEnvironments DeletePrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/AppServiceEnvironments_DeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: Deletes a private endpoint connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesAPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .deletePrivateEndpointConnection(
                "test-rg", "test-ase", "fa38656c-034e-43d8-adce-fe06ce039c98", Context.NONE);
    }
}
