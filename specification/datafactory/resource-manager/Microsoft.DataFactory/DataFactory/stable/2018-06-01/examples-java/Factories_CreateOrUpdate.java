
/**
 * Samples for Factories CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Factories_CreateOrUpdate.json
     */
    /**
     * Sample code: Factories_CreateOrUpdate.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesCreateOrUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.factories().define("exampleFactoryName").withExistingResourceGroup("exampleResourceGroup")
            .withRegion("East US").create();
    }
}
