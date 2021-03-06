import com.azure.core.util.Context;

/** Samples for DataVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/DataVersionBase/delete.json
     */
    /**
     * Sample code: Delete Data Version Base.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteDataVersionBase(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.dataVersions().deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
