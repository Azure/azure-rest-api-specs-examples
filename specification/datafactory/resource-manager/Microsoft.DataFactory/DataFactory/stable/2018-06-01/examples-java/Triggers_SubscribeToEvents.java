
/**
 * Samples for Triggers SubscribeToEvents.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Triggers_SubscribeToEvents.json
     */
    /**
     * Sample code: Triggers_SubscribeToEvents.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersSubscribeToEvents(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggers().subscribeToEvents("exampleResourceGroup", "exampleFactoryName", "exampleTrigger",
            com.azure.core.util.Context.NONE);
    }
}
