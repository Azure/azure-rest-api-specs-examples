import com.azure.core.util.Context;

/** Samples for PrivateDnsZoneGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/PrivateEndpointDnsZoneGroupGet.json
     */
    /**
     * Sample code: Get private dns zone group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateDnsZoneGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateDnsZoneGroups()
            .getWithResponse("rg1", "testPe", "testPdnsgroup", Context.NONE);
    }
}
