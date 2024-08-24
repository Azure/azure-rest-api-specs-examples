
import com.azure.resourcemanager.dns.models.RecordType;

/**
 * Samples for RecordSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetPTRRecordset.json
     */
    /**
     * Sample code: Get PTR recordset.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPTRRecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getRecordSets().getWithResponse("rg1", "0.0.127.in-addr.arpa", "1",
            RecordType.PTR, com.azure.core.util.Context.NONE);
    }
}
