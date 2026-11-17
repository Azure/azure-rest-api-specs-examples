
import com.azure.resourcemanager.eventhubs.fluent.models.EHNamespaceInner;
import com.azure.resourcemanager.eventhubs.models.GeoDRRoleType;
import com.azure.resourcemanager.eventhubs.models.GeoDataReplicationProperties;
import com.azure.resourcemanager.eventhubs.models.NamespaceReplicaLocation;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Namespaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/NamespaceWithGeoDRCreate.json
     */
    /**
     * Sample code: NamespaceWithGeoDRCreate.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void namespaceWithGeoDRCreate(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().createOrUpdate("ResurceGroupSample", "NamespaceGeoDRCreateSample",
            new EHNamespaceInner().withLocation("East US")
                .withGeoDataReplication(new GeoDataReplicationProperties().withMaxReplicationLagDurationInSeconds(60)
                    .withLocations(Arrays.asList(
                        new NamespaceReplicaLocation().withLocationName("eastus").withRoleType(GeoDRRoleType.PRIMARY),
                        new NamespaceReplicaLocation().withLocationName("westus").withRoleType(GeoDRRoleType.SECONDARY),
                        new NamespaceReplicaLocation().withLocationName("centralus")
                            .withRoleType(GeoDRRoleType.SECONDARY)))),
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
