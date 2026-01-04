
import com.azure.resourcemanager.containerregistry.fluent.models.TaskInner;
import com.azure.resourcemanager.containerregistry.models.TaskStatus;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Tasks Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/TasksCreate_QuickTask.json
     */
    /**
     * Sample code: Tasks_Create_QuickTask.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tasksCreateQuickTask(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTasks().create("myResourceGroup", "myRegistry",
            "quicktask",
            new TaskInner().withLocation("eastus").withTags(mapOf("testkey", "fakeTokenPlaceholder"))
                .withStatus(TaskStatus.ENABLED).withLogTemplate("acr/tasks:{{.Run.OS}}").withIsSystemTask(true),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
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
