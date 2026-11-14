
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetAAAADelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone AAAA Record Set.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        dELETEPrivateDNSZoneAAAARecordSet(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().deleteWithResponse("resourceGroup1", "privatezone1.com",
            RecordType.AAAA, "recordAAAA", null, com.azure.core.util.Context.NONE);
    }
}
