import com.azure.core.util.Context;

/** Samples for NetworkProfiles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkProfileDelete.json
     */
    /**
     * Sample code: Delete network profile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkProfiles().delete("rg1", "networkProfile1", Context.NONE);
    }
}
