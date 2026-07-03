
/**
 * Samples for Subnets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SubnetList.json
     */
    /**
     * Sample code: List subnets.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listSubnets(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubnets().list("subnet-test", "vnetname", com.azure.core.util.Context.NONE);
    }
}
