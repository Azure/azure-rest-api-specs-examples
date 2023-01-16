/** Samples for DelegatedNetwork ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/controllerListByRG.json
     */
    /**
     * Sample code: Get DelegatedNetwork resources by resource group.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getDelegatedNetworkResourcesByResourceGroup(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.delegatedNetworks().listByResourceGroup("testRG", com.azure.core.util.Context.NONE);
    }
}
