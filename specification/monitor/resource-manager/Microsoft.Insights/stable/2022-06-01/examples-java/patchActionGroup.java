
import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.models.ActionGroupPatchBody;
import java.util.HashMap;
import java.util.Map;

/** Samples for ActionGroups Update. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/patchActionGroup.json
     */
    /**
     * Sample code: Patch an action group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchAnActionGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActionGroups().updateWithResponse(
            "Default-NotificationRules", "SampleActionGroup",
            new ActionGroupPatchBody().withTags(mapOf("key1", "value1", "key2", "value2")).withEnabled(false),
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
