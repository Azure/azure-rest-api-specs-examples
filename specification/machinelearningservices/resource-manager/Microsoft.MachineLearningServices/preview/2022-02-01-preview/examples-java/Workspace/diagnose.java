import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.DiagnoseRequestProperties;
import com.azure.resourcemanager.machinelearning.models.DiagnoseWorkspaceParameters;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workspaces Diagnose. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/diagnose.json
     */
    /**
     * Sample code: Diagnose Workspace.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void diagnoseWorkspace(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .workspaces()
            .diagnose(
                "workspace-1234",
                "testworkspace",
                new DiagnoseWorkspaceParameters()
                    .withValue(
                        new DiagnoseRequestProperties()
                            .withUdr(mapOf())
                            .withNsg(mapOf())
                            .withResourceLock(mapOf())
                            .withDnsResolution(mapOf())
                            .withStorageAccount(mapOf())
                            .withKeyVault(mapOf())
                            .withContainerRegistry(mapOf())
                            .withApplicationInsights(mapOf())
                            .withOthers(mapOf())),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
