
/**
 * Samples for GlobalParameters DeleteSync.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * GlobalParameters_Delete.json
     */
    /**
     * Sample code: GlobalParameters_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void globalParametersDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.globalParameters().deleteWithResponse("exampleResourceGroup", "exampleFactoryName", "default",
            com.azure.core.util.Context.NONE);
    }
}
