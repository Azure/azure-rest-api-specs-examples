import com.azure.core.util.Context;

/** Samples for GlobalParameters ListByFactory. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_ListByFactory.json
     */
    /**
     * Sample code: GlobalParameters_ListByFactory.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void globalParametersListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.globalParameters().listByFactory("exampleResourceGroup", "exampleFactoryName", Context.NONE);
    }
}
