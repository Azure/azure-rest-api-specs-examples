import com.azure.resourcemanager.dns.models.RecordType;

/** Samples for RecordSets ListByType. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListNSRecordset.json
     */
    /**
     * Sample code: List NS recordsets.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNSRecordsets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .listByType("rg1", "zone1", RecordType.NS, null, null, com.azure.core.util.Context.NONE);
    }
}
