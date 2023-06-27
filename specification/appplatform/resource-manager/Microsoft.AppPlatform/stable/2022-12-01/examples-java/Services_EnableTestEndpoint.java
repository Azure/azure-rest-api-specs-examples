import com.azure.core.util.Context;

/** Samples for Services EnableTestEndpoint. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Services_EnableTestEndpoint.json
     */
    /**
     * Sample code: Services_EnableTestEndpoint.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesEnableTestEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServices()
            .enableTestEndpointWithResponse("myResourceGroup", "myservice", Context.NONE);
    }
}
