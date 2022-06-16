import com.azure.core.util.Context;

/** Samples for Triggers Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Stop.json
     */
    /**
     * Sample code: Triggers_Stop.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersStop(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggers().stop("exampleResourceGroup", "exampleFactoryName", "exampleTrigger", Context.NONE);
    }
}
