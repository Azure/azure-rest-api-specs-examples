
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.privatedns.fluent.models.RecordSetInner;
import com.azure.resourcemanager.privatedns.models.RecordType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RecordSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/PTMLinkRecordSetAPut.json
     */
    /**
     * Sample code: PUT Private DNS Zone PTMLink A Record Set.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        pUTPrivateDNSZonePTMLinkARecordSet(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().createOrUpdateWithResponse("resourceGroup1", "privatezone1.com",
            RecordType.A, "ptmlinkrecord",
            new RecordSetInner().withMetadata(mapOf("key1", "fakeTokenPlaceholder")).withTtl(3600L)
                .withTrafficManagementProfile(new SubResource().withId(
                    "/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateTrafficManagerProfiles/ptm-profile1")),
            null, null, com.azure.core.util.Context.NONE);
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
