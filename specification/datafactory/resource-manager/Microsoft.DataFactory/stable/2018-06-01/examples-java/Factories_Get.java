
/**
 * Samples for Factories GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_Get.json
     */
    /**
     * Sample code: Factories_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.factories().getByResourceGroupWithResponse("exampleResourceGroup", "exampleFactoryName", null,
            com.azure.core.util.Context.NONE);
    }
}
