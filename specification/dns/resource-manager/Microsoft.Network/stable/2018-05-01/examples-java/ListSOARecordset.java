import com.azure.resourcemanager.dns.models.RecordType;

/** Samples for RecordSets ListByType. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListSOARecordset.json
     */
    /**
     * Sample code: List SOA recordsets.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSOARecordsets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .listByType("rg1", "zone1", RecordType.SOA, null, null, com.azure.core.util.Context.NONE);
    }
}
