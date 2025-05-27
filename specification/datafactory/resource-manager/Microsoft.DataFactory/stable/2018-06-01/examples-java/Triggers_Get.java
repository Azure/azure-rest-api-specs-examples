
/**
 * Samples for Triggers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Get.json
     */
    /**
     * Sample code: Triggers_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggers().getWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleTrigger", null,
            com.azure.core.util.Context.NONE);
    }
}
