
/** Samples for Zones Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/DeleteZone.json
     */
    /**
     * Sample code: Delete zone.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getZones().delete("rg1", "zone1", null,
            com.azure.core.util.Context.NONE);
    }
}
