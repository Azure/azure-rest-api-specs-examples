
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineExtensionInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachineExtensions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineExamples/VirtualMachineExtension_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_CreateOrUpdate_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineExtensionCreateOrUpdateMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineExtensions().createOrUpdate("rgcompute",
            "myVM", "myVMExtension", new VirtualMachineExtensionInner().withLocation("westus"),
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
