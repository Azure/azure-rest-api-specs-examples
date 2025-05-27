
/**
 * Samples for PrivateEndpointConnectionOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * GetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Get a private endpoint connection for a datafactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        getAPrivateEndpointConnectionForADatafactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.privateEndpointConnectionOperations().getWithResponse("exampleResourceGroup", "exampleFactoryName",
            "connection", null, com.azure.core.util.Context.NONE);
    }
}
