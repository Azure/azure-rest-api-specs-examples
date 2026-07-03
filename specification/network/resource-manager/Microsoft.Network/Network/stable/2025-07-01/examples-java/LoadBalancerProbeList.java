
/**
 * Samples for LoadBalancerProbes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerProbeList.json
     */
    /**
     * Sample code: LoadBalancerProbeList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void loadBalancerProbeList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerProbes().list("testrg", "lb", com.azure.core.util.Context.NONE);
    }
}
