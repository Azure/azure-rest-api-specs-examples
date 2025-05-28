
/**
 * Samples for TriggerRuns Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/TriggerRuns_Cancel.
     * json
     */
    /**
     * Sample code: Triggers_Cancel.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersCancel(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggerRuns().cancelWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleTrigger",
            "2f7fdb90-5df1-4b8e-ac2f-064cfa58202b", com.azure.core.util.Context.NONE);
    }
}
