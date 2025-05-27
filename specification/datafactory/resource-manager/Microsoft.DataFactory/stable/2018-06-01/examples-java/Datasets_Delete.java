
/**
 * Samples for Datasets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Delete.json
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
