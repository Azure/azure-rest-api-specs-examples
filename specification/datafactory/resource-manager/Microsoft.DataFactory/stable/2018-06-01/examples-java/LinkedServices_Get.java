
/**
 * Samples for LinkedServices Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/LinkedServices_Get.
     * json
     */
    /**
     * Sample code: LinkedServices_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void linkedServicesGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.linkedServices().getWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleLinkedService",
            null, com.azure.core.util.Context.NONE);
    }
}
