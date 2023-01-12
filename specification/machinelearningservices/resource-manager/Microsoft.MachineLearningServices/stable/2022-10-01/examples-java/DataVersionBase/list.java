/** Samples for DataVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/DataVersionBase/list.json
     */
    /**
     * Sample code: List Data Version Base.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listDataVersionBase(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .dataVersions()
            .list(
                "test-rg",
                "my-aml-workspace",
                "string",
                "string",
                1,
                null,
                "string",
                null,
                com.azure.core.util.Context.NONE);
    }
}
