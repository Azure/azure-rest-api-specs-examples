
import com.azure.resourcemanager.chaos.models.ActionDependency;
import com.azure.resourcemanager.chaos.models.ActionDependencyType;
import com.azure.resourcemanager.chaos.models.ActionLifecycle;
import com.azure.resourcemanager.chaos.models.FailureBehavior;
import com.azure.resourcemanager.chaos.models.RunAfter;
import com.azure.resourcemanager.chaos.models.RunAfterBehavior;
import com.azure.resourcemanager.chaos.models.ScenarioAction;
import com.azure.resourcemanager.chaos.models.ScenarioProperties;
import java.util.Arrays;

/**
 * Samples for Scenarios CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Scenarios_CreateOrUpdate_With_Failure_Behavior.json
     */
    /**
     * Sample code: Create or update a scenario that continues past suppressed action failures.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createOrUpdateAScenarioThatContinuesPastSuppressedActionFailures(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarios().define("bestEffortProbeScenario").withExistingWorkspace("exampleRG", "exampleWorkspace")
            .withProperties(new ScenarioProperties()
                .withDescription(
                    "Probes dependencies best-effort, then runs a critical shutdown whose failure must stop the run.")
                .withActions(Arrays.asList(new ScenarioAction().withName("dependencyProbe")
                    .withActionId("urn:csci:microsoft:agent:networkDisconnect/1.0.0")
                    .withDescription(
                        "Best-effort dependency probe. Inherits the scenario-level Continue behavior, so a failure here does not fail the run.")
                    .withDuration("PT5M"),
                    new ScenarioAction().withName("criticalShutdown")
                        .withActionId("urn:csci:microsoft:compute:shutdown/1.0.0")
                        .withDescription(
                            "Critical shutdown. Overrides the scenario default so that a failure stops the run.")
                        .withDuration("PT10M")
                        .withRunAfter(new RunAfter().withBehavior(RunAfterBehavior.ALL)
                            .withItems(Arrays.asList(new ActionDependency().withType(ActionDependencyType.ACTION)
                                .withName("dependencyProbe").withOnActionLifecycle(ActionLifecycle.ANY_TERMINAL))))
                        .withFailureBehavior(FailureBehavior.STOP)))
                .withFailureBehavior(FailureBehavior.CONTINUE))
            .create();
    }
}
