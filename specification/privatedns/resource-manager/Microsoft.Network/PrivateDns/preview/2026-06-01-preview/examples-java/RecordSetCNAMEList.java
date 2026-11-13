
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets ListByType.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetCNAMEList.json
     */
    /**
     * Sample code: GET Private DNS Zone CNAME Record Sets.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        gETPrivateDNSZoneCNAMERecordSets(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().listByType("resourceGroup1", "privatezone1.com", RecordType.CNAME, null,
            null, com.azure.core.util.Context.NONE);
    }
}
