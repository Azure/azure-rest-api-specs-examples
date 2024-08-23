
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/operationsList.json
     */
    /**
     * Sample code: OperationsList.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void operationsList(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
