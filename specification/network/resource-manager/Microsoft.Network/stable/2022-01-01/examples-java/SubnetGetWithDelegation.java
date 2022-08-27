import com.azure.core.util.Context;

/** Samples for Subnets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/SubnetGetWithDelegation.json
     */
    /**
     * Sample code: Get subnet with a delegation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSubnetWithADelegation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSubnets()
            .getWithResponse("subnet-test", "vnetname", "subnet1", null, Context.NONE);
    }
}
