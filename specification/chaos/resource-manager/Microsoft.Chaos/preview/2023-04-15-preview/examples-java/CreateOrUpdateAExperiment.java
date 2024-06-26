import com.azure.resourcemanager.chaos.models.Branch;
import com.azure.resourcemanager.chaos.models.ContinuousAction;
import com.azure.resourcemanager.chaos.models.KeyValuePair;
import com.azure.resourcemanager.chaos.models.ListSelector;
import com.azure.resourcemanager.chaos.models.ResourceIdentity;
import com.azure.resourcemanager.chaos.models.ResourceIdentityType;
import com.azure.resourcemanager.chaos.models.Step;
import com.azure.resourcemanager.chaos.models.TargetReference;
import com.azure.resourcemanager.chaos.models.TargetReferenceType;
import java.time.Duration;
import java.util.Arrays;

/** Samples for Experiments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/CreateOrUpdateAExperiment.json
     */
    /**
     * Sample code: Create/update a Experiment in a resource group.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void createUpdateAExperimentInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager
            .experiments()
            .define("exampleExperiment")
            .withRegion("eastus2euap")
            .withExistingResourceGroup("exampleRG")
            .withSteps(
                Arrays
                    .asList(
                        new Step()
                            .withName("step1")
                            .withBranches(
                                Arrays
                                    .asList(
                                        new Branch()
                                            .withName("branch1")
                                            .withActions(
                                                Arrays
                                                    .asList(
                                                        new ContinuousAction()
                                                            .withName("urn:csci:microsoft:virtualMachine:shutdown/1.0")
                                                            .withDuration(Duration.parse("PT10M"))
                                                            .withParameters(
                                                                Arrays
                                                                    .asList(
                                                                        new KeyValuePair()
                                                                            .withKey("fakeTokenPlaceholder")
                                                                            .withValue("false")))
                                                            .withSelectorId("selector1")))))))
            .withSelectors(
                Arrays
                    .asList(
                        new ListSelector()
                            .withId("selector1")
                            .withTargets(
                                Arrays
                                    .asList(
                                        new TargetReference()
                                            .withType(TargetReferenceType.CHAOS_TARGET)
                                            .withId(
                                                "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine")))))
            .withIdentity(new ResourceIdentity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .create();
    }
}
