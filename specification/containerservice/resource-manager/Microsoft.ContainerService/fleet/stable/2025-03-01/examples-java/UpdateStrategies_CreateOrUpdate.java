
import com.azure.resourcemanager.containerservicefleet.models.UpdateGroup;
import com.azure.resourcemanager.containerservicefleet.models.UpdateRunStrategy;
import com.azure.resourcemanager.containerservicefleet.models.UpdateStage;
import java.util.Arrays;

/**
 * Samples for FleetUpdateStrategies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/UpdateStrategies_CreateOrUpdate.json
     */
    /**
     * Sample code: Create a FleetUpdateStrategy.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void createAFleetUpdateStrategy(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetUpdateStrategies().define("strartegy1").withExistingFleet("rg1", "fleet1")
            .withStrategy(new UpdateRunStrategy().withStages(Arrays.asList(new UpdateStage().withName("stage1")
                .withGroups(Arrays.asList(new UpdateGroup().withName("group-a"))).withAfterStageWaitInSeconds(3600))))
            .create();
    }
}
