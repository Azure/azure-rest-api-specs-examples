
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetPTRDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone PTR Record Set.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        dELETEPrivateDNSZonePTRRecordSet(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().deleteWithResponse("resourceGroup1", "0.0.127.in-addr.arpa",
            RecordType.PTR, "1", null, com.azure.core.util.Context.NONE);
    }
}
