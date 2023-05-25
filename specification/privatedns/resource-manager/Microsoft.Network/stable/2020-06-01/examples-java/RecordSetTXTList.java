import com.azure.resourcemanager.privatedns.models.RecordType;

/** Samples for RecordSets ListByType. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetTXTList.json
     */
    /**
     * Sample code: GET Private DNS Zone TXT Record Sets.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETPrivateDNSZoneTXTRecordSets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .privateDnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .listByType(
                "resourceGroup1", "privatezone1.com", RecordType.TXT, null, null, com.azure.core.util.Context.NONE);
    }
}
