
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetAGet.json
     */
    /**
     * Sample code: GET Private DNS Zone A Record Set.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void gETPrivateDNSZoneARecordSet(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().getWithResponse("resourceGroup1", "privatezone1.com", RecordType.A,
            "recordA", com.azure.core.util.Context.NONE);
    }
}
