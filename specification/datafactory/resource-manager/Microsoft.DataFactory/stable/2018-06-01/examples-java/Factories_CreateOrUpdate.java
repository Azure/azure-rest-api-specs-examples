
/**
 * Samples for Factories CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * Factories_CreateOrUpdate.json
     */
    /**
     * Sample code: Factories_CreateOrUpdate.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesCreateOrUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.factories().define("exampleFactoryName").withRegion("East US")
            .withExistingResourceGroup("exampleResourceGroup").create();
    }
}
