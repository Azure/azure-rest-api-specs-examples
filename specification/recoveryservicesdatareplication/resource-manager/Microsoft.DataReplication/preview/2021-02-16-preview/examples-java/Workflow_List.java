/** Samples for Workflow List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Workflow_List.json
     */
    /**
     * Sample code: Workflow_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void workflowList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .workflows()
            .list(
                "rgrecoveryservicesdatareplication",
                "4",
                "mnebpgmjcitjleipnttx",
                "rdavrzbethhslmkqgajontnxsue",
                com.azure.core.util.Context.NONE);
    }
}
