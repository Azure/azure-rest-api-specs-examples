import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/operationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void operationsList(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.operations().list(Context.NONE);
    }
}
