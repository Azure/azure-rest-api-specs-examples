
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetPTRGet.json
     */
    /**
     * Sample code: GET Private DNS Zone PTR Record Set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETPrivateDNSZonePTRRecordSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getRecordSets().getWithResponse("resourceGroup1",
            "0.0.127.in-addr.arpa", RecordType.PTR, "1", com.azure.core.util.Context.NONE);
    }
}
