import com.azure.resourcemanager.recoveryservicessiterecovery.models.CreateProtectionContainerInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ReplicationProviderSpecificContainerCreationInput;
import java.util.Arrays;

/** Samples for ReplicationProtectionContainers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationProtectionContainers_Create.json
     */
    /**
     * Sample code: Create a protection container.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void createAProtectionContainer(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationProtectionContainers()
            .define("cloud_6d224fc6-f326-5d35-96de-fbf51efb3179")
            .withExistingReplicationFabric("vault1", "resourceGroupPS1", "cloud1")
            .withProperties(
                new CreateProtectionContainerInputProperties()
                    .withProviderSpecificInput(Arrays.asList(new ReplicationProviderSpecificContainerCreationInput())))
            .create();
    }
}
