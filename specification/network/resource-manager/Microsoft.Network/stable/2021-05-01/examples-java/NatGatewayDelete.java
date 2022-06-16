import com.azure.core.util.Context;

/** Samples for NatGateways Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NatGatewayDelete.json
     */
    /**
     * Sample code: Delete nat gateway.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNatGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNatGateways().delete("rg1", "test-natGateway", Context.NONE);
    }
}
