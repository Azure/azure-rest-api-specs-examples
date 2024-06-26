/** Samples for DataVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/DataVersionBase/delete.json
     */
    /**
     * Sample code: Delete Data Version Base.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteDataVersionBase(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .dataVersions()
            .deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", com.azure.core.util.Context.NONE);
    }
}
