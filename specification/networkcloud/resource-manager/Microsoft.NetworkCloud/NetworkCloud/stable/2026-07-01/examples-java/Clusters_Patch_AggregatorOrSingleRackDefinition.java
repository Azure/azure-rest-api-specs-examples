
import com.azure.resourcemanager.networkcloud.models.AdministrativeCredentialsPatch;
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineConfigurationDataPatch;
import com.azure.resourcemanager.networkcloud.models.Cluster;
import com.azure.resourcemanager.networkcloud.models.RackDefinitionPatch;
import com.azure.resourcemanager.networkcloud.models.StorageApplianceConfigurationDataPatch;
import com.azure.resourcemanager.networkcloud.models.ValidationThresholdGrouping;
import com.azure.resourcemanager.networkcloud.models.ValidationThresholdPatch;
import com.azure.resourcemanager.networkcloud.models.ValidationThresholdType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Clusters_Patch_AggregatorOrSingleRackDefinition.json
     */
    /**
     * Sample code: Patch cluster AggregatorOrSingleRackDefinition.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchClusterAggregatorOrSingleRackDefinition(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        Cluster resource = manager.clusters()
            .getByResourceGroupWithResponse("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withAggregatorOrSingleRackDefinition(new RackDefinitionPatch()
                .withBareMetalMachineConfigurationData(Arrays.asList(
                    new BareMetalMachineConfigurationDataPatch()
                        .withBmcCredentials(new AdministrativeCredentialsPatch().withPassword("fakeTokenPlaceholder")
                            .withUsername("username"))
                        .withBmcMacAddress("AA:BB:CC:DD:EE:FF").withBootMacAddress("00:BB:CC:DD:EE:FF")
                        .withMachineDetails("extraDetails").withMachineName("bmmName1").withRackSlot(1L)
                        .withSerialNumber("BM1219XXX"),
                    new BareMetalMachineConfigurationDataPatch()
                        .withBmcCredentials(new AdministrativeCredentialsPatch().withPassword("fakeTokenPlaceholder")
                            .withUsername("username"))
                        .withBmcMacAddress("AA:BB:CC:DD:EE:00").withBootMacAddress("00:BB:CC:DD:EE:00")
                        .withMachineDetails("extraDetails").withMachineName("bmmName2").withRackSlot(2L)
                        .withSerialNumber("BM1219YYY")))
                .withNetworkRackId(
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName")
                .withRackLocation("Foo Datacenter, Floor 3, Aisle 9, Rack 2").withRackSerialNumber("newSerialNumber")
                .withRackSkuId(
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName")
                .withStorageApplianceConfigurationData(Arrays.asList(new StorageApplianceConfigurationDataPatch()
                    .withAdminCredentials(new AdministrativeCredentialsPatch().withPassword("fakeTokenPlaceholder")
                        .withUsername("username"))
                    .withRackSlot(1L).withSerialNumber("BM1219XXX").withStorageApplianceName("vmName"))))
            .withComputeDeploymentThreshold(
                new ValidationThresholdPatch().withGrouping(ValidationThresholdGrouping.PER_CLUSTER)
                    .withType(ValidationThresholdType.PERCENT_SUCCESS).withValue(90L))
            .apply();
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
