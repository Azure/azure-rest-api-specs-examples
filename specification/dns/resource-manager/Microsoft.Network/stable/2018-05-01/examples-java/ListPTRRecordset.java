import com.azure.resourcemanager.dns.models.RecordType;

/** Samples for RecordSets ListByType. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListPTRRecordset.json
     */
    /**
     * Sample code: List PTR recordsets.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPTRRecordsets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .listByType("rg1", "0.0.127.in-addr.arpa", RecordType.PTR, null, null, com.azure.core.util.Context.NONE);
    }
}
