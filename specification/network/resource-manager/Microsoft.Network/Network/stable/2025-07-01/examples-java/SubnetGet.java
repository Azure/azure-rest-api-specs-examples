
/**
 * Samples for Subnets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SubnetGet.json
     */
    /**
     * Sample code: Get subnet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getSubnet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubnets().getWithResponse("subnet-test", "vnetname", "subnet1", null,
            com.azure.core.util.Context.NONE);
    }
}
