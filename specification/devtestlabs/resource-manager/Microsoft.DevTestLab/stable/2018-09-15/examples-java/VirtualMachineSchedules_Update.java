
import com.azure.resourcemanager.devtestlabs.models.ScheduleFragment;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachineSchedules Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/
     * VirtualMachineSchedules_Update.json
     */
    /**
     * Sample code: VirtualMachineSchedules_Update.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachineSchedulesUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachineSchedules().updateWithResponse("resourceGroupName", "{labName}", "{vmName}",
            "LabVmsShutdown", new ScheduleFragment().withTags(mapOf("tagName1", "tagValue1")),
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
