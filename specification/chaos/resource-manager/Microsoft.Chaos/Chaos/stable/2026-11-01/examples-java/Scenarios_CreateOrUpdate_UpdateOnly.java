
import com.azure.resourcemanager.chaos.models.Scenario;
import com.azure.resourcemanager.chaos.models.ScenarioAction;
import com.azure.resourcemanager.chaos.models.ScenarioActionSelector;
import com.azure.resourcemanager.chaos.models.ScenarioActionSelectorName;
import com.azure.resourcemanager.chaos.models.ScenarioProperties;
import java.util.Arrays;

/**
 * Samples for Scenarios CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Scenarios_CreateOrUpdate_UpdateOnly.json
     */
    /**
     * Sample code: Update a scenario only when present, or return 412 when it does not exist.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void updateAScenarioOnlyWhenPresentOrReturn412WhenItDoesNotExist(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        Scenario resource = manager.scenarios()
            .getWithResponse("exampleRG", "exampleWorkspace", "exampleScenario", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new ScenarioProperties().withParameters(Arrays.asList())
            .withActions(Arrays.asList(new ScenarioAction().withName("restart")
                .withActionId("microsoft-compute-shutdown/2.0").withDuration("PT5M")
                .withSelectors(Arrays.asList(
                    new ScenarioActionSelector().withName(ScenarioActionSelectorName.LOGICAL_ZONE).withValue("1"),
                    new ScenarioActionSelector().withName(ScenarioActionSelectorName.LOCATION).withValue("eastus"))))))
            .withIfMatch("*").apply();
    }
}
