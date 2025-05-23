
import com.azure.resourcemanager.compute.fluent.models.DedicatedHostInner;
import com.azure.resourcemanager.compute.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DedicatedHosts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * dedicatedHostExamples/DedicatedHost_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a dedicated host .
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateADedicatedHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHosts().createOrUpdate("myResourceGroup",
            "myDedicatedHostGroup", "myDedicatedHost",
            new DedicatedHostInner().withLocation("westus").withTags(mapOf("department", "HR"))
                .withSku(new Sku().withName("DSv3-Type1")).withPlatformFaultDomain(1),
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
