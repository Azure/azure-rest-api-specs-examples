import com.azure.core.util.Context;

/** Samples for CodeVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/CodeVersion/list.json
     */
    /**
     * Sample code: List Code Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeVersions().list("test-rg", "my-aml-workspace", "string", "string", 1, null, Context.NONE);
    }
}
