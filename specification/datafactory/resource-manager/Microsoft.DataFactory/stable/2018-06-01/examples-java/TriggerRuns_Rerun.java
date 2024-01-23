
/**
 * Samples for TriggerRuns Rerun.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/TriggerRuns_Rerun.
     * json
     */
    /**
     * Sample code: Triggers_Rerun.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersRerun(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggerRuns().rerunWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleTrigger",
            "2f7fdb90-5df1-4b8e-ac2f-064cfa58202b", com.azure.core.util.Context.NONE);
    }
}
