import com.azure.core.util.Context;

/** Samples for PrivateZones Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/examples/PrivateZoneDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dELETEPrivateDNSZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .privateDnsZones()
            .manager()
            .serviceClient()
            .getPrivateZones()
            .delete("resourceGroup1", "privatezone1.com", null, Context.NONE);
    }
}
