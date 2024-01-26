
import com.azure.resourcemanager.dns.models.RecordType;

/** Samples for RecordSets Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetNSRecordset.json
     */
    /**
     * Sample code: Get NS recordset.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getNSRecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getRecordSets().getWithResponse("rg1", "zone1", "record1",
            RecordType.NS, com.azure.core.util.Context.NONE);
    }
}
