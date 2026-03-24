
/**
 * Samples for Factories Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Factories_Delete.json
     */
    /**
     * Sample code: Factories_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.factories().deleteByResourceGroupWithResponse("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
