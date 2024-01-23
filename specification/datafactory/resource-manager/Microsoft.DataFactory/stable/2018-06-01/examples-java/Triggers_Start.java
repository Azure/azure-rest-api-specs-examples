
/**
 * Samples for Triggers Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Start.json
     */
    /**
     * Sample code: Triggers_Start.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersStart(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggers().start("exampleResourceGroup", "exampleFactoryName", "exampleTrigger",
            com.azure.core.util.Context.NONE);
    }
}
