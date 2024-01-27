
import com.azure.resourcemanager.privatedns.models.RecordType;

/** Samples for RecordSets ListByType. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetPTRList.json
     */
    /**
     * Sample code: GET Private DNS Zone PTR Record Sets.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETPrivateDNSZonePTRRecordSets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getRecordSets().listByType("resourceGroup1",
            "0.0.127.in-addr.arpa", RecordType.PTR, null, null, com.azure.core.util.Context.NONE);
    }
}
