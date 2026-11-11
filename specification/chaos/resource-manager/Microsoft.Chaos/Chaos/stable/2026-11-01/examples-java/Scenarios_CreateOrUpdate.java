
import com.azure.resourcemanager.chaos.models.ExternalResource;
import com.azure.resourcemanager.chaos.models.KeyValuePair;
import com.azure.resourcemanager.chaos.models.ParameterType;
import com.azure.resourcemanager.chaos.models.ScenarioAction;
import com.azure.resourcemanager.chaos.models.ScenarioActionSelector;
import com.azure.resourcemanager.chaos.models.ScenarioActionSelectorName;
import com.azure.resourcemanager.chaos.models.ScenarioParameter;
import com.azure.resourcemanager.chaos.models.ScenarioProperties;
import java.util.Arrays;

/**
 * Samples for Scenarios CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Scenarios_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a scenario.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createOrUpdateAScenario(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarios().define("zoneDownScenario").withExistingWorkspace("exampleRG", "exampleWorkspace")
            .withProperties(new ScenarioProperties().withDescription(
                "Induces an outage of all discovered VM and VMSS instances in the target zone with an option to invoke custom scripted actions using Automation Runbooks. Additionally, it forces failover of discovered Redis instances to simulate backend zonal outage.")
                .withParameters(Arrays.asList(
                    new ScenarioParameter().withName("logicalZone").withType(ParameterType.STRING).withRequired(true)
                        .withDescription(
                            "Logical availability-zone label to test in each target subscription, for example 1."),
                    new ScenarioParameter().withName("location").withType(ParameterType.STRING).withRequired(true)
                        .withDescription("Azure location name in which to test the logical zone, for example eastus."),
                    new ScenarioParameter().withName("duration").withType(ParameterType.STRING)
                        .withDefaultProperty("PT15M").withRequired(false)
                        .withDescription("The duration of the outage scenario."),
                    new ScenarioParameter().withName("customRunbook1ResourceId").withType(ParameterType.STRING)
                        .withRequired(false).withDescription(
                            "Optional custom runbook 1 resource ID. If not provided, this action will be skipped."),
                    new ScenarioParameter().withName("customRunbook1ParametersJson").withType(ParameterType.STRING)
                        .withDefaultProperty("{}").withRequired(false)
                        .withDescription("Optional custom runbook 1 parameters in JSON format.")))
                .withActions(Arrays.asList(
                    new ScenarioAction().withName("vmssZoneDown").withActionId("microsoft-compute-shutdown/2.0")
                        .withDescription("Force shutdown VMSS instances in target zone")
                        .withDuration("%%{parameters.duration}%%").withParameters(Arrays.asList())
                        .withSelectors(Arrays.asList(
                            new ScenarioActionSelector().withName(ScenarioActionSelectorName.LOGICAL_ZONE)
                                .withValue("%%{parameters.logicalZone}%%"),
                            new ScenarioActionSelector().withName(ScenarioActionSelectorName.LOCATION)
                                .withValue("%%{parameters.location}%%"))),
                    new ScenarioAction().withName("vmZoneDown").withActionId("microsoft-compute-shutdown/2.0")
                        .withDescription("Force shutdown VM instances in target zone")
                        .withDuration("%%{parameters.duration}%%").withParameters(Arrays.asList())
                        .withSelectors(Arrays.asList(
                            new ScenarioActionSelector().withName(ScenarioActionSelectorName.LOGICAL_ZONE)
                                .withValue("%%{parameters.logicalZone}%%"),
                            new ScenarioActionSelector().withName(ScenarioActionSelectorName.LOCATION)
                                .withValue("%%{parameters.location}%%"))),
                    new ScenarioAction().withName("redisFailover")
                        .withActionId("urn:csci:microsoft:azureClusteredCacheForRedis:Reboot/1.0.0")
                        .withDescription("Force failover of Redis instances to simulate backend zonal outage")
                        .withDuration("PT5M").withParameters(
                            Arrays.asList(new KeyValuePair().withKey("fakeTokenPlaceholder").withValue("PrimaryNode"))),
                    new ScenarioAction().withName("custom-runbook-1")
                        .withActionId("urn:csci:microsoft:Automation:StartRunbook/1.0.0")
                        .withDescription("Custom Runbook 1").withDuration("PT30M")
                        .withParameters(Arrays.asList(new KeyValuePair().withKey("fakeTokenPlaceholder")
                            .withValue("%%{parameters.customRunbook1ParametersJson}%%")))
                        .withExternalResource(new ExternalResource().withResourceId(
                            "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Automation/automationAccounts/exampleAutomationAccount/runbooks/exampleRunbook")))))
            .create();
    }
}
