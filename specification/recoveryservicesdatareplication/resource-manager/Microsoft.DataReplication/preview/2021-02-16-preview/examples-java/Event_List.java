/** Samples for Event List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Event_List.json
     */
    /**
     * Sample code: Event_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void eventList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .events()
            .list(
                "rgrecoveryservicesdatareplication",
                "4",
                "wbglupjzvkirtgnnyasxom",
                "cxtufi",
                com.azure.core.util.Context.NONE);
    }
}
