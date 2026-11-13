
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets ListByType.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetPTRList.json
     */
    /**
     * Sample code: GET Private DNS Zone PTR Record Sets.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        gETPrivateDNSZonePTRRecordSets(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().listByType("resourceGroup1", "0.0.127.in-addr.arpa", RecordType.PTR,
            null, null, com.azure.core.util.Context.NONE);
    }
}
