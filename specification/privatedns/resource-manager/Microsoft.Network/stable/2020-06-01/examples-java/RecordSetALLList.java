/** Samples for RecordSets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetALLList.json
     */
    /**
     * Sample code: GET Private DNS Zone ALL Record Sets.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETPrivateDNSZoneALLRecordSets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .privateDnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .list("resourceGroup1", "privatezone1.com", null, null, com.azure.core.util.Context.NONE);
    }
}
