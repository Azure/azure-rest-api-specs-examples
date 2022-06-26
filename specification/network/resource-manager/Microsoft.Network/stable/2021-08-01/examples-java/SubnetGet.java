import com.azure.core.util.Context;

/** Samples for Subnets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/SubnetGet.json
     */
    /**
     * Sample code: Get subnet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSubnet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSubnets()
            .getWithResponse("subnet-test", "vnetname", "subnet1", null, Context.NONE);
    }
}
