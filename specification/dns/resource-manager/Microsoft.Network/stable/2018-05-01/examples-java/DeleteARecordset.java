import com.azure.resourcemanager.dns.models.RecordType;

/** Samples for RecordSets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/DeleteARecordset.json
     */
    /**
     * Sample code: Delete CNAME recordset.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteCNAMERecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .deleteWithResponse("rg1", "zone1", "record1", RecordType.A, null, com.azure.core.util.Context.NONE);
    }
}
