/** Samples for ReplicationNetworks ListByReplicationFabrics. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationNetworks_ListByReplicationFabrics.json
     */
    /**
     * Sample code: Gets the list of networks under a fabric.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfNetworksUnderAFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationNetworks()
            .listByReplicationFabrics(
                "srce2avaultbvtaC27",
                "srcBvte2a14C27",
                "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac",
                com.azure.core.util.Context.NONE);
    }
}
