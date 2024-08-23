
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetTXTDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone TXT Record Set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dELETEPrivateDNSZoneTXTRecordSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getRecordSets().deleteWithResponse("resourceGroup1",
            "privatezone1.com", RecordType.TXT, "recordTXT", null, com.azure.core.util.Context.NONE);
    }
}
