
/**
 * Samples for ReplicationNetworks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationNetworks_Get.json
     */
    /**
     * Sample code: Gets a network with specified server id and network name.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsANetworkWithSpecifiedServerIdAndNetworkName(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationNetworks().getWithResponse("srcBvte2a14C27", "srce2avaultbvtaC27",
            "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac", "93ce99d7-1219-4914-aa61-73fe5023988e",
            com.azure.core.util.Context.NONE);
    }
}
