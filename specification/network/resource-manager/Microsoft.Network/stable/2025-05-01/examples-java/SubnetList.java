
/**
 * Samples for Subnets List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/SubnetList.json
     */
    /**
     * Sample code: List subnets.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSubnets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSubnets().list("subnet-test", "vnetname",
            com.azure.core.util.Context.NONE);
    }
}
