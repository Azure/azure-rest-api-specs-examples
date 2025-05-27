
/**
 * Samples for DataFlowDebugSession QueryByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * DataFlowDebugSession_QueryByFactory.json
     */
    /**
     * Sample code: DataFlowDebugSession_QueryByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        dataFlowDebugSessionQueryByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.dataFlowDebugSessions().queryByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
