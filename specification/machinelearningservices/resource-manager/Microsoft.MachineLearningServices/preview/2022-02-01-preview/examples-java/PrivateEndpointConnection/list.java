import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/PrivateEndpointConnection/list.json
     */
    /**
     * Sample code: StorageAccountListPrivateEndpointConnections.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void storageAccountListPrivateEndpointConnections(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.privateEndpointConnections().list("rg-1234", "testworkspace", Context.NONE);
    }
}
