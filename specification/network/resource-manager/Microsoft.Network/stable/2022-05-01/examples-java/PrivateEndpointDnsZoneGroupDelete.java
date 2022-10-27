import com.azure.core.util.Context;

/** Samples for PrivateDnsZoneGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/PrivateEndpointDnsZoneGroupDelete.json
     */
    /**
     * Sample code: Delete private dns zone group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePrivateDnsZoneGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateDnsZoneGroups()
            .delete("rg1", "testPe", "testPdnsgroup", Context.NONE);
    }
}
