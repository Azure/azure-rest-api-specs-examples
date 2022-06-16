import com.azure.core.util.Context;

/** Samples for DataVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/DataVersionBase/get.json
     */
    /**
     * Sample code: Get Data Version Base.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getDataVersionBase(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.dataVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
