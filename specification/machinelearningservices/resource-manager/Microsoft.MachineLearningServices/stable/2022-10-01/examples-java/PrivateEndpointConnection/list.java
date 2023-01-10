import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/PrivateEndpointConnection/list.json
     */
    /**
     * Sample code: StorageAccountListPrivateEndpointConnections.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void storageAccountListPrivateEndpointConnections(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.privateEndpointConnections().list("rg-1234", "testworkspace", Context.NONE);
    }
}
