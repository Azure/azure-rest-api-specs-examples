
import com.azure.resourcemanager.chaos.models.ConfigurationExclusions;
import com.azure.resourcemanager.chaos.models.ConfigurationFilters;
import com.azure.resourcemanager.chaos.models.KeyValuePair;
import com.azure.resourcemanager.chaos.models.ScenarioConfigurationProperties;
import java.util.Arrays;

/**
 * Samples for ScenarioConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ScenarioConfigurations_CreateOrUpdate_With_Physical_Zones.json
     */
    /**
     * Sample code: Create or update a scenario configuration with physical zone targeting.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createOrUpdateAScenarioConfigurationWithPhysicalZoneTargeting(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioConfigurations().define("config-physical-zone")
            .withExistingScenario("exampleRG", "exampleWorkspace", "12345678-1234-1234-1234-123456789012")
            .withProperties(new ScenarioConfigurationProperties().withScenarioId(
                "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012")
                .withParameters(Arrays.asList(new KeyValuePair().withKey("fakeTokenPlaceholder").withValue("PT10M")))
                .withExclusions(new ConfigurationExclusions().withResources(Arrays.asList(
                    "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/protectedVM")))
                .withFilters(new ConfigurationFilters().withLocations(Arrays.asList("westus2"))
                    .withPhysicalZones(Arrays.asList("westus2-az1"))))
            .create();
    }
}
