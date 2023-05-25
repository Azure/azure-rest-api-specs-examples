/** Samples for Zones GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetZone.json
     */
    /**
     * Sample code: Get zone.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getZones()
            .getByResourceGroupWithResponse("rg1", "zone1", com.azure.core.util.Context.NONE);
    }
}
