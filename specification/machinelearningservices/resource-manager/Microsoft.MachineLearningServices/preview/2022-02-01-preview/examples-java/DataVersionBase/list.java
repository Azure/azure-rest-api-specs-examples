import com.azure.core.util.Context;

/** Samples for DataVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/DataVersionBase/list.json
     */
    /**
     * Sample code: List Data Version Base.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listDataVersionBase(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .dataVersions()
            .list("test-rg", "my-aml-workspace", "string", "string", 1, null, "string", null, Context.NONE);
    }
}
