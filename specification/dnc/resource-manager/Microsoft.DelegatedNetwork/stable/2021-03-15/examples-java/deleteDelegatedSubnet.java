/** Samples for DelegatedSubnetService Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/deleteDelegatedSubnet.json
     */
    /**
     * Sample code: delete delegated subnet.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void deleteDelegatedSubnet(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.delegatedSubnetServices().delete("TestRG", "delegated1", null, com.azure.core.util.Context.NONE);
    }
}
