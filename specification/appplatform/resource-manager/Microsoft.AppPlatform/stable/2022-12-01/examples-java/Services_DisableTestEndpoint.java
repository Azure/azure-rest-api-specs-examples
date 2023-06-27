import com.azure.core.util.Context;

/** Samples for Services DisableTestEndpoint. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Services_DisableTestEndpoint.json
     */
    /**
     * Sample code: Services_DisableTestEndpoint.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesDisableTestEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServices()
            .disableTestEndpointWithResponse("myResourceGroup", "myservice", Context.NONE);
    }
}
