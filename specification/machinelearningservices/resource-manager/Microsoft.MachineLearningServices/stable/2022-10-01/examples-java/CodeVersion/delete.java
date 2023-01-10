import com.azure.core.util.Context;

/** Samples for CodeVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/CodeVersion/delete.json
     */
    /**
     * Sample code: Delete Code Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeVersions().deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
