
import com.azure.resourcemanager.containerservicefleet.models.GateConfiguration;
import com.azure.resourcemanager.containerservicefleet.models.GateType;
import com.azure.resourcemanager.containerservicefleet.models.ManagedClusterUpdate;
import com.azure.resourcemanager.containerservicefleet.models.ManagedClusterUpgradeSpec;
import com.azure.resourcemanager.containerservicefleet.models.ManagedClusterUpgradeType;
import com.azure.resourcemanager.containerservicefleet.models.NodeImageSelection;
import com.azure.resourcemanager.containerservicefleet.models.NodeImageSelectionType;
import com.azure.resourcemanager.containerservicefleet.models.UpdateGroup;
import com.azure.resourcemanager.containerservicefleet.models.UpdateRunStrategy;
import com.azure.resourcemanager.containerservicefleet.models.UpdateStage;
import java.util.Arrays;

/**
 * Samples for UpdateRuns CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/UpdateRuns_CreateOrUpdate.json
     */
    /**
     * Sample code: Create an UpdateRun.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        createAnUpdateRun(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.updateRuns().define("run1").withExistingFleet("rg1", "fleet1").withUpdateStrategyId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateStrategies/strategy1")
            .withStrategy(new UpdateRunStrategy().withStages(Arrays.asList(new UpdateStage().withName("stage1")
                .withGroups(Arrays.asList(new UpdateGroup().withName("group-a")
                    .withBeforeGates(Arrays.asList(
                        new GateConfiguration().withDisplayName("gate before group-a").withType(GateType.APPROVAL)))
                    .withAfterGates(Arrays.asList(
                        new GateConfiguration().withDisplayName("gate after group-a").withType(GateType.APPROVAL)))))
                .withAfterStageWaitInSeconds(3600)
                .withBeforeGates(Arrays
                    .asList(new GateConfiguration().withDisplayName("gate before stage1").withType(GateType.APPROVAL)))
                .withAfterGates(Arrays.asList(
                    new GateConfiguration().withDisplayName("gate after stage1").withType(GateType.APPROVAL))))))
            .withManagedClusterUpdate(new ManagedClusterUpdate()
                .withUpgrade(new ManagedClusterUpgradeSpec().withType(ManagedClusterUpgradeType.FULL)
                    .withKubernetesVersion("1.26.1"))
                .withNodeImageSelection(new NodeImageSelection().withType(NodeImageSelectionType.LATEST)))
            .create();
    }
}
