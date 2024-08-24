
import com.azure.resourcemanager.dns.models.RecordType;

/**
 * Samples for RecordSets ListByType.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListCaaRecordset.json
     */
    /**
     * Sample code: List CAA recordsets.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCAARecordsets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getRecordSets().listByType("rg1", "zone1", RecordType.CAA, null,
            null, com.azure.core.util.Context.NONE);
    }
}
