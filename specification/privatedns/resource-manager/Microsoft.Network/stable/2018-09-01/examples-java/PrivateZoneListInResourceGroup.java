import com.azure.core.util.Context;

/** Samples for PrivateZones ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/examples/PrivateZoneListInResourceGroup.json
     */
    /**
     * Sample code: GET Private DNS Zone by Resource Group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETPrivateDNSZoneByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .privateDnsZones()
            .manager()
            .serviceClient()
            .getPrivateZones()
            .listByResourceGroup("resourceGroup1", null, Context.NONE);
    }
}
