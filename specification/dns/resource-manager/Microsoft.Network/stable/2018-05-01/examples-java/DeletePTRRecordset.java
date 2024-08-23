
import com.azure.resourcemanager.dns.models.RecordType;

/**
 * Samples for RecordSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/DeletePTRRecordset.json
     */
    /**
     * Sample code: Delete PTR recordset.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePTRRecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getRecordSets().deleteWithResponse("rg1", "0.0.127.in-addr.arpa",
            "1", RecordType.PTR, null, com.azure.core.util.Context.NONE);
    }
}
