
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachineScaleSets Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Update_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetUpdateMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().update("rgcompute", "aaaaaaaaaaaaaa",
            new VirtualMachineScaleSetUpdate(), null, null, com.azure.core.util.Context.NONE);
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
