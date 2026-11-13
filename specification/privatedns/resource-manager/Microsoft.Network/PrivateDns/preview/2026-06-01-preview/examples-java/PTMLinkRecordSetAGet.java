
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/PTMLinkRecordSetAGet.json
     */
    /**
     * Sample code: GET Private DNS Zone PTMLink A Record Set.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        gETPrivateDNSZonePTMLinkARecordSet(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().getWithResponse("resourceGroup1", "privatezone1.com", RecordType.A,
            "ptmlinkrecord", com.azure.core.util.Context.NONE);
    }
}
