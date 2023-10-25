import com.azure.resourcemanager.recoveryservicesdatareplication.models.ReplicationExtensionModelCustomProperties;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.ReplicationExtensionModelProperties;

/** Samples for ReplicationExtension Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ReplicationExtension_Create.json
     */
    /**
     * Sample code: ReplicationExtension_Create.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void replicationExtensionCreate(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .replicationExtensions()
            .define("g16yjJ")
            .withExistingReplicationVault("rgrecoveryservicesdatareplication", "4")
            .withProperties(
                new ReplicationExtensionModelProperties()
                    .withCustomProperties(new ReplicationExtensionModelCustomProperties()))
            .create();
    }
}
