/** Samples for Applications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationListOperation_example.json
     */
    /**
     * Sample code: Get a list of application resources.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getAListOfApplicationResources(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applications().listWithResponse("resRg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
