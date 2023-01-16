/** Samples for DelegatedSubnetService ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/delegatedSubnetListByRG.json
     */
    /**
     * Sample code: Get DelegatedSubnets resources by resource group.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getDelegatedSubnetsResourcesByResourceGroup(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.delegatedSubnetServices().listByResourceGroup("testRG", com.azure.core.util.Context.NONE);
    }
}
