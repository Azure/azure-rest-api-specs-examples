/** Samples for EmailConfiguration List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/EmailConfiguration_List.json
     */
    /**
     * Sample code: EmailConfiguration_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void emailConfigurationList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.emailConfigurations().list("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE);
    }
}
