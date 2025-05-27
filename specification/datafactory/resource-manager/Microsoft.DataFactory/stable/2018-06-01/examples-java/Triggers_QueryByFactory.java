
import com.azure.resourcemanager.datafactory.models.TriggerFilterParameters;

/**
 * Samples for Triggers QueryByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * Triggers_QueryByFactory.json
     */
    /**
     * Sample code: Triggers_QueryByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersQueryByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggers().queryByFactoryWithResponse("exampleResourceGroup", "exampleFactoryName",
            new TriggerFilterParameters().withParentTriggerName("exampleTrigger"), com.azure.core.util.Context.NONE);
    }
}
