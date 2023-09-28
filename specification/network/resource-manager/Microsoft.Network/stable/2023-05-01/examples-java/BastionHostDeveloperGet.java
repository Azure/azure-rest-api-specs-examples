/** Samples for BastionHosts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/BastionHostDeveloperGet.json
     */
    /**
     * Sample code: Get Developer Bastion Host.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDeveloperBastionHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getBastionHosts()
            .getByResourceGroupWithResponse("rg1", "bastionhostdeveloper'", com.azure.core.util.Context.NONE);
    }
}
