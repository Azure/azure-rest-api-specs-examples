import com.azure.resourcemanager.dns.models.RecordType;

/** Samples for RecordSets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetSOARecordset.json
     */
    /**
     * Sample code: Get SOA recordset.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSOARecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .getWithResponse("rg1", "zone1", "@", RecordType.SOA, com.azure.core.util.Context.NONE);
    }
}
