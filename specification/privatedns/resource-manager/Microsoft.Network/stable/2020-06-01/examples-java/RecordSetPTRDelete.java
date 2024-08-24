
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetPTRDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone PTR Record Set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dELETEPrivateDNSZonePTRRecordSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getRecordSets().deleteWithResponse("resourceGroup1",
            "0.0.127.in-addr.arpa", RecordType.PTR, "1", null, com.azure.core.util.Context.NONE);
    }
}
