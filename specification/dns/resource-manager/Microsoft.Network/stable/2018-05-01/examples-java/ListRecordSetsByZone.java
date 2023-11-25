/** Samples for RecordSets ListAllByDnsZone. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListRecordSetsByZone.json
     */
    /**
     * Sample code: List recordsets by zone.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRecordsetsByZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .listAllByDnsZone("rg1", "zone1", null, null, com.azure.core.util.Context.NONE);
    }
}
