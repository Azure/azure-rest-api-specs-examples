/** Samples for DelegatedSubnetService GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/getDelegatedSubnet.json
     */
    /**
     * Sample code: Get details of a delegated subnet.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getDetailsOfADelegatedSubnet(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager
            .delegatedSubnetServices()
            .getByResourceGroupWithResponse("TestRG", "delegated1", com.azure.core.util.Context.NONE);
    }
}
