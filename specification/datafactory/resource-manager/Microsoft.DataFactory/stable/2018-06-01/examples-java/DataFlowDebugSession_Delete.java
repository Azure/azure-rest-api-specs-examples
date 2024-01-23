
import com.azure.resourcemanager.datafactory.models.DeleteDataFlowDebugSessionRequest;

/**
 * Samples for DataFlowDebugSession Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * DataFlowDebugSession_Delete.json
     */
    /**
     * Sample code: DataFlowDebugSession_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void dataFlowDebugSessionDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.dataFlowDebugSessions().deleteWithResponse("exampleResourceGroup", "exampleFactoryName",
            new DeleteDataFlowDebugSessionRequest().withSessionId("91fb57e0-8292-47be-89ff-c8f2d2bb2a7e"),
            com.azure.core.util.Context.NONE);
    }
}
