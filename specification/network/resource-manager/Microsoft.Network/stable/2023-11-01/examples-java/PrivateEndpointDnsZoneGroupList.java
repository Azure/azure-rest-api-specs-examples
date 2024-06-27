
/**
 * Samples for PrivateDnsZoneGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/
     * PrivateEndpointDnsZoneGroupList.json
     */
    /**
     * Sample code: List private endpoints in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPrivateEndpointsInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateDnsZoneGroups().list("testPe", "rg1",
            com.azure.core.util.Context.NONE);
    }
}
