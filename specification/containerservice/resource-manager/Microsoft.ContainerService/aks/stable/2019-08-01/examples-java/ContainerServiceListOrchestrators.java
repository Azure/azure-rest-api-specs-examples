
/**
 * Samples for ContainerServices ListOrchestrators.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2019-08-01/examples/
     * ContainerServiceListOrchestrators.json
     */
    /**
     * Sample code: List Container Service Orchestrators.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listContainerServiceOrchestrators(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getContainerServices()
            .listOrchestratorsWithResponse("location1", null, com.azure.core.util.Context.NONE);
    }
}
