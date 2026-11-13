
import com.azure.resourcemanager.containerinstance.fluent.models.ContainerGroupInner;
import com.azure.resourcemanager.containerinstance.models.ConfigMap;
import com.azure.resourcemanager.containerinstance.models.Container;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupProfileReferenceDefinition;
import com.azure.resourcemanager.containerinstance.models.StandbyPoolProfileDefinition;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ContainerGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupCreateOrUpdateStandbyPool.json
     */
    /**
     * Sample code: ContainerGroupCreateOrUpdateWithStandbyPool.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupCreateOrUpdateWithStandbyPool(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainerGroups().createOrUpdate("demo", "demo1", new ContainerGroupInner()
            .withLocation("west us")
            .withContainers(Arrays.asList(new Container().withName("demo1")
                .withConfigMap(new ConfigMap().withKeyValuePairs(mapOf("Newkey", "fakeTokenPlaceholder")))))
            .withContainerGroupProfile(new ContainerGroupProfileReferenceDefinition().withId(
                "/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroupProfiles/democgp")
                .withRevision(1))
            .withStandbyPoolProfile(new StandbyPoolProfileDefinition().withId(
                "/subscriptions/subid/resourceGroups/demo/providers/Microsoft.StandbyPool/standbyContainerGroupPools/demopool")),
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
