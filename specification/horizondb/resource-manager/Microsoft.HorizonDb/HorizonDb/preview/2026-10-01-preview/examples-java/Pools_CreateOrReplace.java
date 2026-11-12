
import com.azure.resourcemanager.horizondb.models.ComputeProcessorType;
import com.azure.resourcemanager.horizondb.models.PoolComputeModel;
import com.azure.resourcemanager.horizondb.models.PoolComputeModelType;
import com.azure.resourcemanager.horizondb.models.PoolHighAvailability;
import com.azure.resourcemanager.horizondb.models.PoolHighAvailabilityMode;
import com.azure.resourcemanager.horizondb.models.PoolProperties;
import com.azure.resourcemanager.horizondb.models.SystemScheduledMaintenanceProperties;

/**
 * Samples for Pools CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Pools_CreateOrReplace.json
     */
    /**
     * Sample code: Create a minimum provisioned HorizonDB pool using the default parameter group.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void createAMinimumProvisionedHorizonDBPoolUsingTheDefaultParameterGroup(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.pools().define("example-pool-01").withExistingCluster("example-resource-group", "example-cluster")
            .withProperties(new PoolProperties()
                .withComputeModel(new PoolComputeModel().withType(PoolComputeModelType.PROVISIONED)
                    .withComputeProcessor(ComputeProcessorType.AMD).withComputeGeneration(3).withVCores(2))
                .withNodeCount(1)
                .withHighAvailability(new PoolHighAvailability().withMode(PoolHighAvailabilityMode.DISABLED))
                .withScheduledMaintenance(new SystemScheduledMaintenanceProperties()))
            .create();
    }
}
