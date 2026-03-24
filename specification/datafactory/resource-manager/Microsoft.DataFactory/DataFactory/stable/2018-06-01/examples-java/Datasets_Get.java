
/**
 * Samples for Datasets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Datasets_Get.json
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
