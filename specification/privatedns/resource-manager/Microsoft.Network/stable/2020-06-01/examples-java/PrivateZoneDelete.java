
/** Samples for PrivateZones Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/PrivateZoneDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dELETEPrivateDNSZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getPrivateZones().delete("resourceGroup1", "privatezone1.com",
            null, com.azure.core.util.Context.NONE);
    }
}
