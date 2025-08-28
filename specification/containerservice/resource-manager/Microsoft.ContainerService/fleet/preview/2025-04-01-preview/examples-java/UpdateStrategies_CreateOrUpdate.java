
import com.azure.resourcemanager.containerservicefleet.models.GateConfiguration;
import com.azure.resourcemanager.containerservicefleet.models.GateType;
import com.azure.resourcemanager.containerservicefleet.models.UpdateGroup;
import com.azure.resourcemanager.containerservicefleet.models.UpdateRunStrategy;
import com.azure.resourcemanager.containerservicefleet.models.UpdateStage;
import java.util.Arrays;

/**
 * Samples for FleetUpdateStrategies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/UpdateStrategies_CreateOrUpdate.json
     */
    /**
     * Sample code: Create a FleetUpdateStrategy.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void createAFleetUpdateStrategy(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetUpdateStrategies().define("strategy1").withExistingFleet("rg1", "fleet1")
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
            .create();
    }
}
