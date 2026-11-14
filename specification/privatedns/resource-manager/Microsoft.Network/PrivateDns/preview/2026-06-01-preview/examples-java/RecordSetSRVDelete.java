
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetSRVDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone SRV Record Set.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        dELETEPrivateDNSZoneSRVRecordSet(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().deleteWithResponse("resourceGroup1", "privatezone1.com", RecordType.SRV,
            "recordSRV", null, com.azure.core.util.Context.NONE);
    }
}
