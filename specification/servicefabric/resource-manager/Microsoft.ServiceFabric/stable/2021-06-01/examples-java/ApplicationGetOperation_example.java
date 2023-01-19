/** Samples for Applications Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationGetOperation_example.json
     */
    /**
     * Sample code: Get an application.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getAnApplication(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applications().getWithResponse("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
