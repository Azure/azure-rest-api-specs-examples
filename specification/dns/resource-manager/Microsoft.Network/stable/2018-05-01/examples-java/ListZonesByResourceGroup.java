
/** Samples for Zones ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListZonesByResourceGroup.json
     */
    /**
     * Sample code: List zones by resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listZonesByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getZones().listByResourceGroup("rg1", null,
            com.azure.core.util.Context.NONE);
    }
}
