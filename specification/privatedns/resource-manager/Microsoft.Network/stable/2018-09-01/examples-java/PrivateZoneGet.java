import com.azure.core.util.Context;

/** Samples for PrivateZones GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/examples/PrivateZoneGet.json
     */
    /**
     * Sample code: GET Private DNS Zone.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETPrivateDNSZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .privateDnsZones()
            .manager()
            .serviceClient()
            .getPrivateZones()
            .getByResourceGroupWithResponse("resourceGroup1", "privatezone1.com", Context.NONE);
    }
}
