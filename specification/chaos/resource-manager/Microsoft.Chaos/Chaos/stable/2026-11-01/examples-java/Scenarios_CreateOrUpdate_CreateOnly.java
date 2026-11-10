
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
     * x-ms-original-file: 2026-11-01/Scenarios_CreateOrUpdate_CreateOnly.json
     */
    /**
     * Sample code: Create a scenario only when absent, or return 412 when it already exists.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createAScenarioOnlyWhenAbsentOrReturn412WhenItAlreadyExists(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarios().define("exampleScenario").withExistingWorkspace("exampleRG", "exampleWorkspace")
            .withProperties(new ScenarioProperties().withParameters(Arrays.asList())
                .withActions(Arrays.asList(new ScenarioAction().withName("restart")
                    .withActionId("microsoft-compute-shutdown/2.0").withDuration("PT5M")
                    .withSelectors(Arrays.asList(
                        new ScenarioActionSelector().withName(ScenarioActionSelectorName.LOGICAL_ZONE).withValue("1"),
                        new ScenarioActionSelector().withName(ScenarioActionSelectorName.LOCATION)
                            .withValue("eastus"))))))
            .withIfNoneMatch("*").create();
    }
}
