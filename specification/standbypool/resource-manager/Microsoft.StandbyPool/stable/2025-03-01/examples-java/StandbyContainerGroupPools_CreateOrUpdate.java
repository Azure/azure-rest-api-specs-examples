
import com.azure.resourcemanager.standbypool.models.ContainerGroupProfile;
import com.azure.resourcemanager.standbypool.models.ContainerGroupProperties;
import com.azure.resourcemanager.standbypool.models.RefillPolicy;
import com.azure.resourcemanager.standbypool.models.StandbyContainerGroupPoolElasticityProfile;
import com.azure.resourcemanager.standbypool.models.StandbyContainerGroupPoolResourceProperties;
import com.azure.resourcemanager.standbypool.models.Subnet;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StandbyContainerGroupPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/StandbyContainerGroupPools_CreateOrUpdate.json
     */
    /**
     * Sample code: StandbyContainerGroupPools_CreateOrUpdate.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void
        standbyContainerGroupPoolsCreateOrUpdate(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyContainerGroupPools().define("pool").withRegion("West US")
            .withExistingResourceGroup("rgstandbypool").withTags(mapOf())
            .withProperties(new StandbyContainerGroupPoolResourceProperties()
                .withElasticityProfile(new StandbyContainerGroupPoolElasticityProfile().withMaxReadyCapacity(688L)
                    .withRefillPolicy(RefillPolicy.ALWAYS))
                .withContainerGroupProperties(new ContainerGroupProperties()
                    .withContainerGroupProfile(new ContainerGroupProfile().withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.ContainerInstance/containerGroupProfiles/cgProfile")
                        .withRevision(1L))
                    .withSubnetIds(Arrays.asList(new Subnet().withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Network/virtualNetworks/cgSubnet/subnets/cgSubnet"))))
                .withZones(Arrays.asList("1", "2", "3")))
            .create();
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
