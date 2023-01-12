/** Samples for CodeVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/CodeVersion/get.json
     */
    /**
     * Sample code: Get Code Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .codeVersions()
            .getWithResponse("test-rg", "my-aml-workspace", "string", "string", com.azure.core.util.Context.NONE);
    }
}
