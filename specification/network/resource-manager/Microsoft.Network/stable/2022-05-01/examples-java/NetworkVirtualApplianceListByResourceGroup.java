import com.azure.core.util.Context;

/** Samples for NetworkVirtualAppliances ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkVirtualApplianceListByResourceGroup.json
     */
    /**
     * Sample code: List all Network Virtual Appliance for a given resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllNetworkVirtualApplianceForAGivenResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkVirtualAppliances()
            .listByResourceGroup("rg1", Context.NONE);
    }
}
