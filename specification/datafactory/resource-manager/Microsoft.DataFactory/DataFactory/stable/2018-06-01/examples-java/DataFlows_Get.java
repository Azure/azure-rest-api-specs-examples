
/**
 * Samples for DataFlows Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/DataFlows_Get.json
     */
    /**
     * Sample code: DataFlows_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void dataFlowsGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.dataFlows().getWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleDataFlow", null,
            com.azure.core.util.Context.NONE);
    }
}
