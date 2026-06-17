
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
     * x-ms-original-file: 2026-05-01-preview/ScenarioConfigurations_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a scenario configuration.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createOrUpdateAScenarioConfiguration(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioConfigurations().define("config-5678-9012-3456-789012345678")
            .withExistingScenario("exampleRG", "exampleWorkspace", "12345678-1234-1234-1234-123456789012")
            .withProperties(new ScenarioConfigurationProperties().withScenarioId(
                "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012")
                .withParameters(Arrays.asList(new KeyValuePair().withKey("fakeTokenPlaceholder").withValue("PT10M"),
                    new KeyValuePair().withKey("fakeTokenPlaceholder").withValue(
                        "[\"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/vm1\",\"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/vm2\"]")))
                .withExclusions(new ConfigurationExclusions().withResources(Arrays.asList(
                    "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/protectedVM"))
                    .withTags(Arrays.asList(new KeyValuePair().withKey("fakeTokenPlaceholder").withValue("production")))
                    .withTypes(Arrays.asList("Microsoft.Compute/virtualMachineScaleSets")))
                .withFilters(
                    new ConfigurationFilters().withLocations(Arrays.asList("eastus")).withZones(Arrays.asList("1"))))
            .create();
    }
}
