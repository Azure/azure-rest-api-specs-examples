import com.azure.resourcemanager.dns.models.RecordType;

/** Samples for RecordSets ListByType. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListMXRecordset.json
     */
    /**
     * Sample code: List MX recordsets.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listMXRecordsets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .listByType("rg1", "zone1", RecordType.MX, null, null, com.azure.core.util.Context.NONE);
    }
}
