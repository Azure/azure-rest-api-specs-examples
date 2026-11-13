
import com.azure.resourcemanager.privatedns.models.RecordType;

/**
 * Samples for RecordSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetPTRGet.json
     */
    /**
     * Sample code: GET Private DNS Zone PTR Record Set.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        gETPrivateDNSZonePTRRecordSet(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().getWithResponse("resourceGroup1", "0.0.127.in-addr.arpa",
            RecordType.PTR, "1", com.azure.core.util.Context.NONE);
    }
}
