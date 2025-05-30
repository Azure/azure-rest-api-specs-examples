
import com.azure.resourcemanager.compute.fluent.models.DedicatedHostGroupInner;
import com.azure.resourcemanager.compute.models.DedicatedHostGroupPropertiesAdditionalCapabilities;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DedicatedHostGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * dedicatedHostExamples/DedicatedHostGroup_CreateOrUpdate_WithUltraSSD.json
     */
    /**
     * Sample code: Create or update a dedicated host group with Ultra SSD support.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateADedicatedHostGroupWithUltraSSDSupport(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHostGroups().createOrUpdateWithResponse(
            "myResourceGroup", "myDedicatedHostGroup",
            new DedicatedHostGroupInner().withLocation("westus").withTags(mapOf("department", "finance"))
                .withZones(Arrays.asList("1")).withPlatformFaultDomainCount(3).withSupportAutomaticPlacement(true)
                .withAdditionalCapabilities(
                    new DedicatedHostGroupPropertiesAdditionalCapabilities().withUltraSsdEnabled(true)),
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
