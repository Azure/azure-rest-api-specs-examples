
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetCNAMEDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone CNAME Record Set.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        dELETEPrivateDNSZoneCNAMERecordSet(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().deleteWithResponse("resourceGroup1", "privatezone1.com",
            RecordType.CNAME, "recordCNAME", null, com.azure.core.util.Context.NONE);
    }
}
