/** Samples for Services List. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServiceListOperation_example.json
     */
    /**
     * Sample code: Get a list of service resources.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getAListOfServiceResources(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.services().listWithResponse("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
