
import com.azure.resourcemanager.privatedns.models.RecordType;

/** Samples for RecordSets Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetSOAGet.json
     */
    /**
     * Sample code: GET Private DNS Zone SOA Record Set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETPrivateDNSZoneSOARecordSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getRecordSets().getWithResponse("resourceGroup1",
            "privatezone1.com", RecordType.SOA, "@", com.azure.core.util.Context.NONE);
    }
}
