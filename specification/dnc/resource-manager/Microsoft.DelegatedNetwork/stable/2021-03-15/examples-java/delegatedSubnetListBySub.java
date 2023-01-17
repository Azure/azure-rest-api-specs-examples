/** Samples for DelegatedSubnetService List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/delegatedSubnetListBySub.json
     */
    /**
     * Sample code: Get DelegatedSubnets resources by subscription.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getDelegatedSubnetsResourcesBySubscription(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.delegatedSubnetServices().list(com.azure.core.util.Context.NONE);
    }
}
