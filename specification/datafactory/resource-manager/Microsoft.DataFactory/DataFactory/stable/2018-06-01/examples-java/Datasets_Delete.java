
/**
 * Samples for Datasets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Datasets_Delete.json
     */
    /**
     * Sample code: Datasets_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void datasetsDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.datasets().deleteWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleDataset",
            com.azure.core.util.Context.NONE);
    }
}
