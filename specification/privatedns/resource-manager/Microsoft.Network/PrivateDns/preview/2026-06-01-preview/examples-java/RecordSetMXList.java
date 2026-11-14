
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets ListByType.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetMXList.json
     */
    /**
     * Sample code: GET Private DNS Zone MX Record Sets.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        gETPrivateDNSZoneMXRecordSets(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().listByType("resourceGroup1", "privatezone1.com", RecordType.MX, null,
            null, com.azure.core.util.Context.NONE);
    }
}
