
/**
 * Samples for AppServiceEnvironments GetInboundNetworkDependenciesEndpoints.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetInboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: Get the network endpoints of all inbound dependencies of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getTheNetworkEndpointsOfAllInboundDependenciesOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getInboundNetworkDependenciesEndpoints(
            "Sample-WestUSResourceGroup", "SampleAse", com.azure.core.util.Context.NONE);
    }
}
