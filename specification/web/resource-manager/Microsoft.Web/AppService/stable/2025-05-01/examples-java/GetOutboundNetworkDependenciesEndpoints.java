
/**
 * Samples for AppServiceEnvironments GetOutboundNetworkDependenciesEndpoints.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetOutboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: Get the network endpoints of all outbound dependencies of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getTheNetworkEndpointsOfAllOutboundDependenciesOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getOutboundNetworkDependenciesEndpoints(
            "Sample-WestUSResourceGroup", "SampleAse", com.azure.core.util.Context.NONE);
    }
}
