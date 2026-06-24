
import com.azure.resourcemanager.compute.fluent.models.InterconnectBlockInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.InterconnectBlockProperties;
import com.azure.resourcemanager.compute.models.Placement;
import com.azure.resourcemanager.compute.models.Sku;
import com.azure.resourcemanager.compute.models.ZonePlacementPolicyType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for InterconnectBlocks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/interconnectBlockExamples/InterconnectBlocks_CreateOrUpdate_AnyZone.json
     */
    /**
     * Sample code: Create or update an Interconnect Block with automatic zone placement.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createOrUpdateAnInterconnectBlockWithAutomaticZonePlacement(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getInterconnectBlocks().createOrUpdate("myResourceGroup", "myInterconnectBlock",
            new InterconnectBlockInner().withLocation("westus").withTags(mapOf("department", "HR"))
                .withProperties(new InterconnectBlockProperties().withInterconnectGroup(new ApiEntityReference().withId(
                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup")))
                .withSku(new Sku().withName("Standard_ND128isr_GB300_v6").withCapacity(18L))
                .withPlacement(new Placement().withZonePlacementPolicy(ZonePlacementPolicyType.ANY)),
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
