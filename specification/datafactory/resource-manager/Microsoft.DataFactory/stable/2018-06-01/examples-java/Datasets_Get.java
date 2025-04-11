
/**
 * Samples for Datasets GetSync.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Get.json
     */
    /**
     * Sample code: Datasets_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void datasetsGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.datasets().getWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleDataset", null,
            com.azure.core.util.Context.NONE);
    }
}
