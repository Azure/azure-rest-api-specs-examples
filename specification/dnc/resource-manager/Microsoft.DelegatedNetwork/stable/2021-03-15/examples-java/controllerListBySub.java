/** Samples for DelegatedNetwork List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/controllerListBySub.json
     */
    /**
     * Sample code: Get DelegatedController resources by subscription.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getDelegatedControllerResourcesBySubscription(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.delegatedNetworks().list(com.azure.core.util.Context.NONE);
    }
}
