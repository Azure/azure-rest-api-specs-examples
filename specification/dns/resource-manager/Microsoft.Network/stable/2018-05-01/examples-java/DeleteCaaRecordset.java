
import com.azure.resourcemanager.dns.models.RecordType;

/**
 * Samples for RecordSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/DeleteCaaRecordset.json
     */
    /**
     * Sample code: Delete CAA recordset.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteCAARecordset(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getRecordSets().deleteWithResponse("rg1", "zone1", "record1",
            RecordType.CAA, null, com.azure.core.util.Context.NONE);
    }
}
