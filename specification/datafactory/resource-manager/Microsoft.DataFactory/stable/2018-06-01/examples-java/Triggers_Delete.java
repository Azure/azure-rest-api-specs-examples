
/**
 * Samples for Triggers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Delete.json
     */
    /**
     * Sample code: Triggers_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggers().deleteWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleTrigger",
            com.azure.core.util.Context.NONE);
    }
}
