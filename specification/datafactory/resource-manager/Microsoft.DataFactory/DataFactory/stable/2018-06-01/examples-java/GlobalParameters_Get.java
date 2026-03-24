
/**
 * Samples for GlobalParameters Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/GlobalParameters_Get.json
     */
    /**
     * Sample code: GlobalParameters_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void globalParametersGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.globalParameters().getWithResponse("exampleResourceGroup", "exampleFactoryName", "default",
            com.azure.core.util.Context.NONE);
    }
}
