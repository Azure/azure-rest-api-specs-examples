import com.azure.resourcemanager.dns.models.RecordType;

/** Samples for RecordSets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetMXRecordset.json
     */
    /**
     * Sample code: Get MX recordset.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getMXRecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .getWithResponse("rg1", "zone1", "record1", RecordType.MX, com.azure.core.util.Context.NONE);
    }
}
