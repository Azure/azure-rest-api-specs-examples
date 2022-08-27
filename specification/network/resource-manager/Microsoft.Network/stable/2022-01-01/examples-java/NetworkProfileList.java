import com.azure.core.util.Context;

/** Samples for NetworkProfiles ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkProfileList.json
     */
    /**
     * Sample code: List resource group network profiles.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listResourceGroupNetworkProfiles(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkProfiles().listByResourceGroup("rg1", Context.NONE);
    }
}
