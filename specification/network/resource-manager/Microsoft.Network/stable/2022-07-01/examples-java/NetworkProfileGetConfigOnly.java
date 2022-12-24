import com.azure.core.util.Context;

/** Samples for NetworkProfiles GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkProfileGetConfigOnly.json
     */
    /**
     * Sample code: Get network profile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getNetworkProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkProfiles()
            .getByResourceGroupWithResponse("rg1", "networkProfile1", null, Context.NONE);
    }
}
