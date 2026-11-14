
/**
 * Samples for RecordSets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RecordSetALLList.json
     */
    /**
     * Sample code: GET Private DNS Zone ALL Record Sets.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        gETPrivateDNSZoneALLRecordSets(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getRecordSets().list("resourceGroup1", "privatezone1.com", null, null,
            com.azure.core.util.Context.NONE);
    }
}
