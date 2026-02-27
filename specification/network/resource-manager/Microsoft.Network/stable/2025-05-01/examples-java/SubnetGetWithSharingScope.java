
/**
 * Samples for Subnets Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/SubnetGetWithSharingScope.
     * json
     */
    /**
     * Sample code: Get subnet with sharing scope.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSubnetWithSharingScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSubnets().getWithResponse("subnet-test", "vnetname", "subnet1",
            null, com.azure.core.util.Context.NONE);
    }
}
