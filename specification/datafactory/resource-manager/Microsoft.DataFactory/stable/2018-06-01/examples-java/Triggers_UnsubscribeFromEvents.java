
/**
 * Samples for Triggers UnsubscribeFromEvents.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * Triggers_UnsubscribeFromEvents.json
     */
    /**
     * Sample code: Triggers_UnsubscribeFromEvents.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersUnsubscribeFromEvents(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggers().unsubscribeFromEvents("exampleResourceGroup", "exampleFactoryName", "exampleTrigger",
            com.azure.core.util.Context.NONE);
    }
}
