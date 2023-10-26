/** Samples for EmailConfiguration Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/EmailConfiguration_Get.json
     */
    /**
     * Sample code: EmailConfiguration_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void emailConfigurationGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .emailConfigurations()
            .getWithResponse("rgrecoveryservicesdatareplication", "4", "0", com.azure.core.util.Context.NONE);
    }
}
