
/**
 * Samples for LinkedServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/LinkedServices_Delete
     * .json
     */
    /**
     * Sample code: LinkedServices_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void linkedServicesDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.linkedServices().deleteWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleLinkedService", com.azure.core.util.Context.NONE);
    }
}
