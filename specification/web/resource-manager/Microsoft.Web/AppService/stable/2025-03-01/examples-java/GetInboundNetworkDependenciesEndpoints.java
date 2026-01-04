
/**
 * Samples for AppServiceEnvironments GetInboundNetworkDependenciesEndpoints.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetInboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: Get the network endpoints of all inbound dependencies of an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheNetworkEndpointsOfAllInboundDependenciesOfAnAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().getInboundNetworkDependenciesEndpoints(
            "Sample-WestUSResourceGroup", "SampleAse", com.azure.core.util.Context.NONE);
    }
}
