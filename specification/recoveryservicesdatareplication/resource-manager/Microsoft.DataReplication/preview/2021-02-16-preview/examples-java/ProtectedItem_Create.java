import com.azure.resourcemanager.recoveryservicesdatareplication.models.ProtectedItemModelCustomProperties;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.ProtectedItemModelProperties;

/** Samples for ProtectedItem Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ProtectedItem_Create.json
     */
    /**
     * Sample code: ProtectedItem_Create.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void protectedItemCreate(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .protectedItems()
            .define("d")
            .withExistingReplicationVault("rgrecoveryservicesdatareplication", "4")
            .withProperties(
                new ProtectedItemModelProperties()
                    .withPolicyName("tjoeiynplt")
                    .withReplicationExtensionName("jwxdo")
                    .withCustomProperties(new ProtectedItemModelCustomProperties()))
            .create();
    }
}
