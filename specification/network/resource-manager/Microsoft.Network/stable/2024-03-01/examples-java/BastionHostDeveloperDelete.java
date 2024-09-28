
/**
 * Samples for BastionHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/BastionHostDeveloperDelete.
     * json
     */
    /**
     * Sample code: Delete Developer Bastion Host.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteDeveloperBastionHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getBastionHosts().delete("rg2", "bastionhostdeveloper",
            com.azure.core.util.Context.NONE);
    }
}
