import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetVMInner;
import java.util.HashMap;
import java.util.Map;

/** Samples for VirtualMachineScaleSetVMs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Update_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMUpdateMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .update(
                "rgcompute",
                "aaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaa",
                new VirtualMachineScaleSetVMInner().withLocation("westus"),
                com.azure.core.util.Context.NONE);
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
