
import com.azure.resourcemanager.chaos.models.KeyValuePair;
import com.azure.resourcemanager.chaos.models.ResourceTargeting;
import com.azure.resourcemanager.chaos.models.ResourceTargetingCriteria;
import com.azure.resourcemanager.chaos.models.ScenarioConfigurationProperties;
import java.util.Arrays;

/**
 * Samples for ScenarioConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ScenarioConfigurations_CreateOrUpdate.json
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
                .withResourceTargeting(new ResourceTargeting()
                    .withInclude(new ResourceTargetingCriteria().withLocations(Arrays.asList("eastus"))
                        .withZones(Arrays.asList("1")))
                    .withExclude(new ResourceTargetingCriteria()
                        .withTypes(Arrays.asList("Microsoft.Compute/virtualMachineScaleSets"))
                        .withTags(
                            Arrays.asList(new KeyValuePair().withKey("fakeTokenPlaceholder").withValue("production")))
                        .withResources(Arrays.asList(
                            "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/protectedVM")))))
            .create();
    }
}
