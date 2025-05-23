
/**
 * Samples for Triggers GetEventSubscriptionStatusSync.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * Triggers_GetEventSubscriptionStatus.json
     */
    /**
     * Sample code: Triggers_GetEventSubscriptionStatus.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        triggersGetEventSubscriptionStatus(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggers().getEventSubscriptionStatusWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleTrigger", com.azure.core.util.Context.NONE);
    }
}
