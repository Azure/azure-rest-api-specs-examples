
import com.azure.resourcemanager.containerregistry.models.TaskStatus;
import com.azure.resourcemanager.containerregistry.models.TaskUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Tasks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/TasksUpdate_QuickTask.json
     */
    /**
     * Sample code: Tasks_Update_QuickTask.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tasksUpdateQuickTask(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTasks().update("myResourceGroup", "myRegistry",
            "quicktask", new TaskUpdateParameters().withTags(mapOf("testkey", "fakeTokenPlaceholder"))
                .withStatus(TaskStatus.ENABLED).withLogTemplate("acr/tasks:{{.Run.OS}}"),
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
