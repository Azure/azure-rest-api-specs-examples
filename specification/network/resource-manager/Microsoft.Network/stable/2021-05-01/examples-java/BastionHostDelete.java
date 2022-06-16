import com.azure.core.util.Context;

/** Samples for BastionHosts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/BastionHostDelete.json
     */
    /**
     * Sample code: Delete Bastion Host.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteBastionHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getBastionHosts().delete("rg1", "bastionhosttenant", Context.NONE);
    }
}
