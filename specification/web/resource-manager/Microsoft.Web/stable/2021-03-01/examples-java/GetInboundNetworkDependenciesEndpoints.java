import com.azure.core.util.Context;

/** Samples for AppServiceEnvironments GetInboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetInboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: Get Inbound Network Dependencies Endpoints.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInboundNetworkDependenciesEndpoints(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .getInboundNetworkDependenciesEndpoints("Sample-WestUSResourceGroup", "SampleAse", Context.NONE);
    }
}
